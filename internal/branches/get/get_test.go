package get

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Indobase/cli/internal/testing/apitest"
	"github.com/Indobase/cli/internal/testing/fstest"
	"github.com/Indobase/cli/internal/utils"
	"github.com/Indobase/cli/internal/utils/flags"
	"github.com/Indobase/cli/pkg/api"
	"github.com/Indobase/cli/pkg/cast"
	"github.com/go-errors/errors"
	"github.com/h2non/gock"
	"github.com/oapi-codegen/nullable"
	"github.com/stretchr/testify/assert"
)

func TestGetBranch(t *testing.T) {
	flags.ProjectRef = apitest.RandomProjectRef()

	t.Run("fetches branch details", func(t *testing.T) {
		t.Cleanup(fstest.MockStdout(t, `
  
   HOST      | PORT | USER   | PASSWORD | JWT SECRET | POSTGRES VERSION             | STATUS         
  -----------|------|--------|----------|------------|------------------------------|----------------
   127.0.0.1 | 5432 | ****** | ******   | ******     | Indobase-postgres-17.4.1.074 | ACTIVE_HEALTHY 

`))
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusOK).
			JSON(api.BranchDetailResponse{
				DbHost:          "127.0.0.1",
				DbPort:          5432,
				PostgresVersion: "Indobase-postgres-17.4.1.074",
				Status:          api.BranchDetailResponseStatusACTIVEHEALTHY,
			})
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.NoError(t, err)
	})

	t.Run("throws error on network error", func(t *testing.T) {
		errNetwork := errors.New("network error")
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			ReplyError(errNetwork)
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.ErrorIs(t, err, errNetwork)
	})

	t.Run("throws error on service unavailable", func(t *testing.T) {
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusServiceUnavailable)
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.ErrorContains(t, err, "unexpected get branch status 503:")
	})
}

func TestTomlOutput(t *testing.T) {
	flags.ProjectRef = apitest.RandomProjectRef()
	// Setup output format
	utils.OutputFormat.Value = utils.OutputToml
	t.Cleanup(func() { utils.OutputFormat.Value = utils.OutputPretty })

	t.Run("encodes toml format", func(t *testing.T) {
		r, w, pipeErr := os.Pipe()
		assert.NoError(t, pipeErr)
		oldStdout := os.Stdout
		os.Stdout = w
		t.Cleanup(func() { os.Stdout = oldStdout })
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusOK).
			JSON(api.BranchDetailResponse{
				DbHost:    "127.0.0.1",
				DbPort:    5432,
				DbUser:    cast.Ptr("postgres"),
				DbPass:    cast.Ptr("postgres"),
				JwtSecret: cast.Ptr("secret-key"),
				Ref:       flags.ProjectRef,
			})
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/api-keys").
			Reply(http.StatusOK).
			JSON([]api.ApiKeyResponse{{
				Name:   "anon",
				ApiKey: nullable.NewNullableWithValue("anon-key"),
			}, {
				Name:   "service_role",
				ApiKey: nullable.NewNullableWithValue("service-role-key"),
			}})
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/config/database/pooler").
			Reply(http.StatusOK).
			JSON([]api.SupavisorConfigResponse{{
				ConnectionString: "postgres://postgres:postgres@127.0.0.1:6543/postgres",
				DatabaseType:     api.SupavisorConfigResponseDatabaseTypePRIMARY,
				PoolMode:         api.SupavisorConfigResponsePoolModeTransaction,
			}})
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.NoError(t, err)
		assert.NoError(t, w.Close())
		data, err := io.ReadAll(r)
		assert.NoError(t, err)
		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		actual := map[string]string{}
		for _, line := range lines {
			parts := strings.SplitN(line, " = ", 2)
			if len(parts) != 2 {
				continue
			}
			actual[parts[0]] = strings.Trim(parts[1], `"`)
		}
		expected := map[string]string{
			"POSTGRES_URL":              "postgresql://postgres:postgres@127.0.0.1:6543/postgres?connect_timeout=10",
			"POSTGRES_URL_NON_POOLING":  "postgresql://postgres:postgres@127.0.0.1:5432/postgres?connect_timeout=10",
			"Indobase_ANON_KEY":         "anon-key",
			"Indobase_JWT_SECRET":       "secret-key",
			"Indobase_SERVICE_ROLE_KEY": "service-role-key",
			"Indobase_URL":              fmt.Sprintf("https://%s.", flags.ProjectRef),
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("throws error on network error", func(t *testing.T) {
		errNetwork := errors.New("network error")
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusOK).
			JSON(api.BranchDetailResponse{
				Ref: flags.ProjectRef,
			})
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/api-keys").
			ReplyError(errNetwork)
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.ErrorIs(t, err, errNetwork)
	})

	t.Run("throws error on database not found", func(t *testing.T) {
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusOK).
			JSON(api.BranchDetailResponse{
				Ref: flags.ProjectRef,
			})
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/api-keys").
			Reply(http.StatusOK).
			JSON([]api.ApiKeyResponse{})
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/config/database/pooler").
			Reply(http.StatusOK).
			JSON([]api.SupavisorConfigResponse{})
		// Run test
		err := Run(context.Background(), flags.ProjectRef, nil)
		assert.ErrorIs(t, err, utils.ErrPrimaryNotFound)
	})
}

func TestBranchDetail(t *testing.T) {
	flags.ProjectRef = apitest.RandomProjectRef()

	t.Run("get branch by name", func(t *testing.T) {
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/branches/main").
			Reply(http.StatusOK).
			JSON(api.BranchResponse{ProjectRef: flags.ProjectRef})
		gock.New(utils.DefaultApiHost).
			Get("/v1/branches/" + flags.ProjectRef).
			Reply(http.StatusOK).
			JSON(api.BranchDetailResponse{})
		// Run test
		_, err := getBranchDetail(context.Background(), "main")
		assert.NoError(t, err)
	})

	t.Run("throws error on network error", func(t *testing.T) {
		errNetwork := errors.New("network error")
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/branches/main").
			ReplyError(errNetwork)
		// Run test
		_, err := getBranchDetail(context.Background(), "main")
		assert.ErrorIs(t, err, errNetwork)
	})

	t.Run("throws error on branch not found", func(t *testing.T) {
		t.Cleanup(apitest.MockPlatformAPI(t))
		// Setup mock api
		gock.New(utils.DefaultApiHost).
			Get("/v1/projects/" + flags.ProjectRef + "/branches/missing").
			Reply(http.StatusNotFound)
		// Run test
		_, err := getBranchDetail(context.Background(), "missing")
		assert.ErrorContains(t, err, "unexpected find branch status 404:")
	})
}
