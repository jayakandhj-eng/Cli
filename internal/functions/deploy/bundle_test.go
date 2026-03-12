package deploy

import (
	"archive/zip"
	"bytes"
	"context"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/Indobase/cli/internal/testing/apitest"
	"github.com/Indobase/cli/internal/utils"
	"github.com/Indobase/cli/pkg/cast"
	functionpkg "github.com/Indobase/cli/pkg/function"
	"github.com/h2non/gock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDockerBundle(t *testing.T) {
	imageUrl := utils.GetRegistryImageUrl(utils.Config.EdgeRuntime.Image)
	utils.EdgeRuntimeId = "test-edge-runtime"
	const containerId = "test-container"

	t.Run("throws error on bundle failure", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		absImportMap := filepath.Join("hello", "deno.json")
		require.NoError(t, utils.WriteFile(absImportMap, []byte("{}"), fsys))
		// Setup deno error
		t.Setenv("TEST_DENO_ERROR", "bundle failed")
		var body bytes.Buffer
		archive := zip.NewWriter(&body)
		w, err := archive.Create("deno")
		require.NoError(t, err)
		_, err = w.Write([]byte("binary"))
		require.NoError(t, err)
		require.NoError(t, archive.Close())
		// Setup mock api
		defer gock.OffAll()
		gock.New("https://github.com").
			Get("/denoland/deno/releases/download/v" + utils.DenoVersion).
			Reply(http.StatusOK).
			Body(&body)
		// Setup mock docker
		require.NoError(t, apitest.MockDocker(utils.Docker))
		apitest.MockDockerStart(utils.Docker, imageUrl, containerId)
		require.NoError(t, apitest.MockDockerLogsExitCode(utils.Docker, containerId, 1))
		// Setup mock bundler
		bundler := NewDockerBundler(fsys)
		// Run test
		meta, err := bundler.Bundle(
			context.Background(),
			"hello",
			filepath.Join("hello", "index.ts"),
			filepath.Join("hello", "deno.json"),
			[]string{filepath.Join("hello", "data.pdf")},
			&body,
		)
		// Check error
		assert.ErrorContains(t, err, "error running container: exit 1")
		assert.Empty(t, apitest.ListUnmatchedRequests())
		expected := functionpkg.NewMetadata(
			"hello",
			filepath.Join("hello", "index.ts"),
			filepath.Join("hello", "deno.json"),
			[]string{filepath.Join("hello", "data.pdf")},
		)
		assert.Equal(t, cast.Ptr("hello"), meta.Name)
		assert.Equal(t, expected.EntrypointPath, meta.EntrypointPath)
		assert.Equal(t, expected.ImportMapPath, meta.ImportMapPath)
		assert.Equal(t, expected.StaticPatterns, meta.StaticPatterns)
		assert.Nil(t, meta.VerifyJwt)
	})

	t.Run("throws error on permission denied", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewReadOnlyFs(afero.NewMemMapFs())
		// Setup mock bundler
		bundler := NewDockerBundler(fsys)
		// Run test
		meta, err := bundler.Bundle(
			context.Background(),
			"hello",
			"hello/index.ts",
			"",
			nil,
			nil,
		)
		// Check error
		assert.ErrorIs(t, err, os.ErrPermission)
		expected := functionpkg.NewMetadata("hello", filepath.Join("hello", "index.ts"), "", nil)
		assert.Equal(t, cast.Ptr("hello"), meta.Name)
		assert.Equal(t, expected.EntrypointPath, meta.EntrypointPath)
		assert.Nil(t, meta.ImportMapPath)
		assert.NotNil(t, meta.StaticPatterns)
		assert.Nil(t, meta.VerifyJwt)
	})
}
