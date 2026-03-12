package format

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/Indobase/cli/internal/utils"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata
var testdata embed.FS

func TestWriteStructured(t *testing.T) {
	testCases, err := testdata.ReadDir("testdata")
	require.NoError(t, err)

	for _, tc := range testCases {
		testName := fmt.Sprintf("formats %s statements", tc.Name())
		sub, err := fs.Sub(testdata, path.Join("testdata", tc.Name()))
		require.NoError(t, err)
		testFs := afero.FromIOFS{FS: sub}
		const dumpPath = "dump.sql"

		t.Run(testName, func(t *testing.T) {
			sql, err := testFs.Open(dumpPath)
			require.NoError(t, err)
			defer sql.Close()
			// Setup in-memory fs
			fsys := afero.NewMemMapFs()
			// Run test
			err = WriteStructuredSchemas(context.Background(), sql, fsys)
			// Check error
			assert.NoError(t, err)
			err = fs.WalkDir(sub, ".", func(fp string, entry fs.DirEntry, err error) error {
				if err != nil || entry.IsDir() || entry.Name() == dumpPath {
					return err
				}
				expected, err := fs.ReadFile(sub, fp)
				assert.NoError(t, err)
				actualPath := filepath.FromSlash(path.Join(utils.IndobaseDirPath, fp))
				actual, _ := afero.ReadFile(fsys, actualPath)
				normalizedExpected := strings.ReplaceAll(string(expected), "\r\n", "\n")
				normalizedActual := strings.ReplaceAll(string(actual), "\r\n", "\n")
				normalizedExpected = strings.TrimRight(normalizedExpected, "\n")
				normalizedActual = strings.TrimRight(normalizedActual, "\n")
				assert.Equal(t, normalizedExpected, normalizedActual, fp)
				return nil
			})
			assert.NoError(t, err)
		})
	}
}

func TestAppendConfig(t *testing.T) {
	t.Run("replaces config inline", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		assert.NoError(t, utils.WriteConfig(fsys, false))
		// Run test
		utils.Config.Db.Migrations.SchemaPaths = []string{
			getSchemaPath("public"),
		}
		err := appendConfig(fsys)
		// Check error
		assert.NoError(t, err)
		data, err := afero.ReadFile(fsys, utils.ConfigPath)
		assert.NoError(t, err)
		var decoded map[string]any
		_, err = toml.Decode(string(data), &decoded)
		assert.NoError(t, err)
		db, ok := decoded["db"].(map[string]any)
		assert.True(t, ok)
		migrations, ok := db["migrations"].(map[string]any)
		assert.True(t, ok)
		pathsRaw, ok := migrations["schema_paths"].([]any)
		assert.True(t, ok)
		var paths []string
		for _, p := range pathsRaw {
			if s, ok := p.(string); ok {
				paths = append(paths, filepath.ToSlash(s))
			}
		}
		assert.Contains(t, paths, "schemas/public/schema.sql")
		normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
		normalized = strings.ReplaceAll(normalized, "\\", "/")
		assert.True(t, strings.HasSuffix(
			strings.TrimSpace(normalized),
			`s3_secret_key = "env(S3_SECRET_KEY)"`,
		))
	})

	t.Run("appends config file", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Run test
		utils.Config.Db.Migrations.SchemaPaths = []string{
			getSchemaPath("public"),
		}
		err := appendConfig(fsys)
		// Check error
		assert.NoError(t, err)
		data, err := afero.ReadFile(fsys, utils.ConfigPath)
		assert.NoError(t, err)
		var decoded map[string]any
		_, err = toml.Decode(string(data), &decoded)
		assert.NoError(t, err)
		db, ok := decoded["db"].(map[string]any)
		assert.True(t, ok)
		migrations, ok := db["migrations"].(map[string]any)
		assert.True(t, ok)
		pathsRaw, ok := migrations["schema_paths"].([]any)
		assert.True(t, ok)
		var paths []string
		for _, p := range pathsRaw {
			if s, ok := p.(string); ok {
				paths = append(paths, filepath.ToSlash(s))
			}
		}
		assert.Contains(t, paths, "schemas/public/schema.sql")
	})
}
