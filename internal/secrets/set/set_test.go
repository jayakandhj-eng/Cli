package set

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/h2non/gock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/Indobase/cli/internal/testing/apitest"
	"github.com/Indobase/cli/internal/utils"
	"github.com/Indobase/cli/pkg/api"
)

func TestSecretSetCommand(t *testing.T) {
	dummy := api.CreateSecretBody{{Name: "my_name", Value: "my_value"}}
	dummyEnv := dummy[0].Name + "=" + dummy[0].Value
	tempDir := filepath.Clean(os.TempDir())
	utils.CurrentDirAbs = tempDir

	t.Run("Sets secret via cli args", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Flush pending mocks after test execution
		defer gock.OffAll()
		gock.New(utils.DefaultApiHost).
			Post("/v1/projects/" + project + "/secrets").
			MatchType("json").
			JSON(dummy).
			Reply(http.StatusCreated)
		// Run test
		err := Run(context.Background(), project, "", []string{dummyEnv}, fsys)
		// Check error
		assert.NoError(t, err)
		assert.Empty(t, apitest.ListUnmatchedRequests())
	})

	t.Run("Sets secret value via env file", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		envPath := filepath.Join(tempDir, ".env")
		require.NoError(t, afero.WriteFile(fsys, envPath, []byte(dummyEnv), 0644))
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Flush pending mocks after test execution
		defer gock.OffAll()
		gock.New(utils.DefaultApiHost).
			Post("/v1/projects/" + project + "/secrets").
			MatchType("json").
			JSON(dummy).
			Reply(http.StatusCreated)
		// Run test
		err := Run(context.Background(), project, envPath, []string{}, fsys)
		// Check error
		assert.NoError(t, err)
		assert.Empty(t, apitest.ListUnmatchedRequests())
	})

	t.Run("throws error on empty secret", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Run test
		err := Run(context.Background(), project, "", []string{}, fsys)
		// Check error
		assert.ErrorContains(t, err, "No arguments found. Use --env-file to read from a .env file.")
	})

	t.Run("throws error on malformed secret", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Run test
		err := Run(context.Background(), project, "", []string{"malformed"}, fsys)
		// Check error
		assert.ErrorContains(t, err, "Invalid secret pair: malformed. Must be NAME=VALUE.")
	})

	t.Run("throws error on network error", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Flush pending mocks after test execution
		defer gock.OffAll()
		gock.New(utils.DefaultApiHost).
			Post("/v1/projects/" + project + "/secrets").
			MatchType("json").
			JSON(dummy).
			ReplyError(errors.New("network error"))
		// Run test
		err := Run(context.Background(), project, "", []string{dummyEnv}, fsys)
		// Check error
		assert.ErrorContains(t, err, "network error")
		assert.Empty(t, apitest.ListUnmatchedRequests())
	})

	t.Run("throws error on server unavailable", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Setup valid project ref
		project := apitest.RandomProjectRef()
		// Setup valid access token
		token := apitest.RandomAccessToken(t)
		t.Setenv("Indobase_ACCESS_TOKEN", string(token))
		// Flush pending mocks after test execution
		defer gock.OffAll()
		gock.New(utils.DefaultApiHost).
			Post("/v1/projects/" + project + "/secrets").
			MatchType("json").
			JSON(dummy).
			Reply(500).
			JSON(map[string]string{"message": "unavailable"})
		// Run test
		err := Run(context.Background(), project, "", []string{dummyEnv}, fsys)
		// Check error
		assert.ErrorContains(t, err, `Unexpected error setting project secrets: {"message":"unavailable"}`)
		assert.Empty(t, apitest.ListUnmatchedRequests())
	})
}

