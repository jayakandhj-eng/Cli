package cmd

import (
	env "github.com/Netflix/go-env"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/Indobase/cli/internal/status"
	"github.com/Indobase/cli/internal/utils"
)

var (
	override []string
	names    status.CustomName

	statusCmd = &cobra.Command{
		GroupID: groupLocalDev,
		Use:     "status",
		Short:   "Show status of local Indobase containers",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			es, err := env.EnvironToEnvSet(override)
			if err != nil {
				return err
			}
			return env.Unmarshal(es, &names)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return status.Run(cmd.Context(), names, utils.OutputFormat.Value, afero.NewOsFs())
		},
		Example: `  Indobase status -o env --override-name api.url=NEXT_PUBLIC_Indobase_URL
  Indobase status -o json`,
	}
)

func init() {
	flags := statusCmd.Flags()
	flags.StringSliceVar(&override, "override-name", []string{}, "Override specific variable names.")
	rootCmd.AddCommand(statusCmd)
}

