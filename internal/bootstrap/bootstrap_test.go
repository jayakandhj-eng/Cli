package bootstrap

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/joho/godotenv"
	"github.com/oapi-codegen/nullable"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/Indobase/cli/internal/utils"
	"github.com/Indobase/cli/internal/utils/flags"
	"github.com/Indobase/cli/pkg/api"
)

func TestSuggestAppStart(t *testing.T) {
	t.Run("suggest npm", func(t *testing.T) {
		cwd, err := os.Getwd()
		require.NoError(t, err)
		// Run test
		suggestion := suggestAppStart(cwd, "npm ci && npm run dev")
		// Check error
		assert.Equal(t, "To start your app:\n  npm ci && npm run dev", suggestion)
	})

	t.Run("suggest cd", func(t *testing.T) {
		cwd, err := os.Getwd()
		require.NoError(t, err)
		// Run test
		suggestion := suggestAppStart(filepath.Dir(cwd), "npm ci && npm run dev")
		// Check error
		expected := "To start your app:"
		expected += "\n  cd " + filepath.Base(cwd)
		expected += "\n  npm ci && npm run dev"
		assert.Equal(t, expected, suggestion)
	})

	t.Run("ignore relative path", func(t *testing.T) {
		// Run test
		suggestion := suggestAppStart(".", "Indobase start")
		// Check error
		assert.Equal(t, "To start your app:\n  Indobase start", suggestion)
	})
}

func TestWriteEnv(t *testing.T) {
	var apiKeys = []api.ApiKeyResponse{{
		Name:   "anon",
		ApiKey: nullable.NewNullableWithValue("anonkey"),
	}, {
		Name:   "service_role",
		ApiKey: nullable.NewNullableWithValue("servicekey"),
	}}

	var dbConfig = pgconn.Config{
		Host:     "db.Indobase.co",
		Port:     5432,
		User:     "admin",
		Password: "password",
		Database: "postgres",
	}

	t.Run("writes .env", func(t *testing.T) {
		flags.ProjectRef = "testing"
		utils.CurrentProfile.ProjectHost = "Indobase.co"
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Run test
		err := writeDotEnv(apiKeys, dbConfig, fsys)
		// Check error
		assert.NoError(t, err)
		env, err := afero.ReadFile(fsys, ".env")
		assert.NoError(t, err)
		envMap, err := godotenv.Parse(strings.NewReader(string(env)))
		assert.NoError(t, err)
		assert.Equal(t, map[string]string{
			POSTGRES_URL:              "postgresql://admin:password@db.Indobase.co:6543/postgres?connect_timeout=10",
			Indobase_ANON_KEY:         "anonkey",
			Indobase_SERVICE_ROLE_KEY: "servicekey",
			Indobase_URL:              "https://testing.Indobase.co",
		}, envMap)
	})

	t.Run("merges with .env.example", func(t *testing.T) {
		flags.ProjectRef = "testing"
		utils.CurrentProfile.ProjectHost = "Indobase.co"
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		example, err := godotenv.Marshal(map[string]string{
			POSTGRES_PRISMA_URL:           "example",
			POSTGRES_URL_NON_POOLING:      "example",
			POSTGRES_USER:                 "example",
			POSTGRES_HOST:                 "example",
			POSTGRES_PASSWORD:             "example",
			POSTGRES_DATABASE:             "example",
			NEXT_PUBLIC_Indobase_ANON_KEY: "example",
			NEXT_PUBLIC_Indobase_URL:      "example",
			"no_match":                    "example",
			Indobase_SERVICE_ROLE_KEY:     "example",
			Indobase_ANON_KEY:             "example",
			Indobase_URL:                  "example",
			POSTGRES_URL:                  "example",
		})
		require.NoError(t, err)
		require.NoError(t, afero.WriteFile(fsys, ".env.example", []byte(example), 0644))
		// Run test
		err = writeDotEnv(apiKeys, dbConfig, fsys)
		// Check error
		assert.NoError(t, err)
		env, err := afero.ReadFile(fsys, ".env")
		assert.NoError(t, err)
		envMap, err := godotenv.Parse(strings.NewReader(string(env)))
		assert.NoError(t, err)
		assert.Equal(t, map[string]string{
			NEXT_PUBLIC_Indobase_ANON_KEY: "anonkey",
			NEXT_PUBLIC_Indobase_URL:      "https://testing.Indobase.co",
			POSTGRES_DATABASE:             "postgres",
			POSTGRES_HOST:                 "db.Indobase.co",
			POSTGRES_PASSWORD:             "password",
			POSTGRES_PRISMA_URL:           "postgresql://admin:password@db.Indobase.co:6543/postgres?connect_timeout=10",
			POSTGRES_URL:                  "postgresql://admin:password@db.Indobase.co:6543/postgres?connect_timeout=10",
			POSTGRES_URL_NON_POOLING:      "postgresql://admin:password@db.Indobase.co:5432/postgres?connect_timeout=10",
			POSTGRES_USER:                 "admin",
			Indobase_ANON_KEY:             "anonkey",
			Indobase_SERVICE_ROLE_KEY:     "servicekey",
			Indobase_URL:                  "https://testing.Indobase.co",
			"no_match":                    "example",
		}, envMap)
	})

	t.Run("throws error on malformed example", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fsys, ".env.example", []byte("!="), 0644))
		// Run test
		err := writeDotEnv(nil, dbConfig, fsys)
		// Check error
		assert.ErrorContains(t, err, `unexpected character "!" in variable name near "!="`)
	})

	t.Run("throws error on permission denied", func(t *testing.T) {
		// Setup in-memory fs
		fsys := afero.NewMemMapFs()
		// Run test
		err := writeDotEnv(nil, dbConfig, afero.NewReadOnlyFs(fsys))
		// Check error
		assert.ErrorIs(t, err, os.ErrPermission)
	})
}

