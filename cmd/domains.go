package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/Indobase/cli/internal/hostnames/activate"
	"github.com/Indobase/cli/internal/hostnames/create"
	"github.com/Indobase/cli/internal/hostnames/delete"
	"github.com/Indobase/cli/internal/hostnames/get"
	"github.com/Indobase/cli/internal/hostnames/reverify"
	"github.com/Indobase/cli/internal/utils"
	"github.com/Indobase/cli/internal/utils/flags"
)

var (
	customHostnamesCmd = &cobra.Command{
		GroupID: groupManagementAPI,
		Use:     "domains",
		Short:   "Manage custom domain names for Indobase projects",
		Long: `Manage custom domain names for Indobase projects.

Use of custom domains and vanity subdomains is mutually exclusive.
`,
	}

	rawOutput      bool
	customHostname string

	customHostnamesCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a custom hostname",
		Long: `Create a custom hostname for your Indobase project.

Expects your custom hostname to have a CNAME record to your Indobase project's subdomain.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if rawOutput && utils.OutputFormat.Value == utils.OutputPretty {
				utils.OutputFormat.Value = utils.OutputJson
			}
			return create.Run(cmd.Context(), flags.ProjectRef, customHostname, afero.NewOsFs())
		},
	}

	customHostnamesGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the current custom hostname config",
		Long:  "Retrieve the custom hostname config for your project, as stored in the Indobase platform.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rawOutput && utils.OutputFormat.Value == utils.OutputPretty {
				utils.OutputFormat.Value = utils.OutputJson
			}
			return get.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}

	customHostnamesReverifyCmd = &cobra.Command{
		Use:   "reverify",
		Short: "Re-verify the custom hostname config for your project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rawOutput && utils.OutputFormat.Value == utils.OutputPretty {
				utils.OutputFormat.Value = utils.OutputJson
			}
			return reverify.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}

	customHostnamesActivateCmd = &cobra.Command{
		Use:   "activate",
		Short: "Activate the custom hostname for a project",
		Long: `Activates the custom hostname configuration for a project.

This reconfigures your Indobase project to respond to requests on your custom hostname.
After the custom hostname is activated, your project's auth services will no longer function on the Indobase-provisioned subdomain.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if rawOutput && utils.OutputFormat.Value == utils.OutputPretty {
				utils.OutputFormat.Value = utils.OutputJson
			}
			return activate.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}

	customHostnamesDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Deletes the custom hostname config for your project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rawOutput && utils.OutputFormat.Value == utils.OutputPretty {
				utils.OutputFormat.Value = utils.OutputJson
			}
			return delete.Run(cmd.Context(), flags.ProjectRef, afero.NewOsFs())
		},
	}
)

func init() {
	persistentFlags := customHostnamesCmd.PersistentFlags()
	persistentFlags.StringVar(&flags.ProjectRef, "project-ref", "", "Project ref of the Indobase project.")
	persistentFlags.BoolVar(&rawOutput, "include-raw-output", false, "Include raw output (useful for debugging).")
	cobra.CheckErr(persistentFlags.MarkDeprecated("include-raw-output", "use -o json instead"))
	createFlags := customHostnamesCreateCmd.Flags()
	createFlags.StringVar(&customHostname, "custom-hostname", "", "The custom hostname to use for your Indobase project.")
	cobra.CheckErr(customHostnamesCreateCmd.MarkFlagRequired("custom-hostname"))
	customHostnamesCmd.AddCommand(customHostnamesGetCmd)
	customHostnamesCmd.AddCommand(customHostnamesCreateCmd)
	customHostnamesCmd.AddCommand(customHostnamesReverifyCmd)
	customHostnamesCmd.AddCommand(customHostnamesActivateCmd)
	customHostnamesCmd.AddCommand(customHostnamesDeleteCmd)
	rootCmd.AddCommand(customHostnamesCmd)
}

