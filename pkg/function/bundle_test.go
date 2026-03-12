package function

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	fs "testing/fstest"

	"github.com/Indobase/cli/pkg/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Setup mock edge runtime binary
	if len(os.Args) > 1 && os.Args[1] == "bundle" {
		if msg := os.Getenv("TEST_BUNDLE_ERROR"); len(msg) > 0 {
			fmt.Fprintln(os.Stderr, msg)
			os.Exit(1)
		}
		os.Exit(0)
	}
	// Run test suite
	os.Exit(m.Run())
}

func TestBundleFunction(t *testing.T) {
	var err error
	edgeRuntimeBin, err = os.Executable()
	require.NoError(t, err)

	t.Run("creates eszip bundle", func(t *testing.T) {
		var body bytes.Buffer
		// Setup in-memory fs
		fsys := fs.MapFS{
			"hello.eszip": &fs.MapFile{},
		}
		// Setup mock bundler
		bundler := nativeBundler{fsys: fsys}
		// Run test
		meta, err := bundler.Bundle(
			context.Background(),
			"hello",
			"hello/index.ts",
			"hello/deno.json",
			[]string{"hello/data.pdf"},
			&body,
		)
		// Check error
		assert.NoError(t, err)
		assert.Equal(t, compressedEszipMagicID+";", body.String())
		assert.Equal(t, cast.Ptr("hello"), meta.Name)
		assert.Equal(t, toFileURL("hello/index.ts"), meta.EntrypointPath)
		importMap := toFileURL("hello/deno.json")
		assert.Equal(t, &importMap, meta.ImportMapPath)
		staticFile := toFileURL("hello/data.pdf")
		assert.Equal(t, cast.Ptr([]string{staticFile}), meta.StaticPatterns)
		assert.Nil(t, meta.VerifyJwt)
	})

	t.Run("ignores empty value", func(t *testing.T) {
		var body bytes.Buffer
		// Setup in-memory fs
		fsys := fs.MapFS{
			"hello.eszip": &fs.MapFile{},
		}
		// Setup mock bundler
		bundler := nativeBundler{fsys: fsys}
		// Run test
		meta, err := bundler.Bundle(
			context.Background(),
			"hello",
			"hello/index.ts",
			"",
			nil,
			&body,
		)
		// Check error
		assert.NoError(t, err)
		assert.Equal(t, compressedEszipMagicID+";", body.String())
		assert.Equal(t, cast.Ptr("hello"), meta.Name)
		assert.Equal(t, toFileURL("hello/index.ts"), meta.EntrypointPath)
		assert.Nil(t, meta.ImportMapPath)
		assert.NotNil(t, meta.StaticPatterns)
		assert.Nil(t, meta.VerifyJwt)
	})
}
