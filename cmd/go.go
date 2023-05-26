package main

import (
	"github.com/coding-wepack/wepack-cli/cmd/require"
	"github.com/coding-wepack/wepack-cli/pkg/cli/golang"
	"github.com/coding-wepack/wepack-cli/pkg/settings"
	"github.com/spf13/cobra"
)

const goHelp = `
This command is used to manager go artifacts.

The command argument must be an arguments, available now:
- push
- pull (TODO)

Common actions for wp:
- wp go push: Publish go artifacts to WePack artifacts repository.

`

func newGoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "go [argument]",
		Short:  "manager go artifacts.",
		Long:   goHelp,
		PreRun: PreRun,
		Args:   require.MinimumNArgs(1),
	}

	// add subcommands
	cmd.AddCommand(
		newGoPushCmd(),
	)

	return cmd
}

const goPushHelp = `
This command publish artifacts to a WePack Artifact Registry.

Examples:
    # Publish go artifacts to a WePack Artifact Registry:
    $ wp go push --module github.com/coding-wepack/wepack-cli@v0.0.1 --repo "https://demo-go.pkg.wepack.net/project/repo/" -u test -p test_pwd

Flags '--module' and '--repo' must be set.
`

func newGoPushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "push",
		Short:  "Publish go artifact to WePack Artifact Repository.",
		Long:   goPushHelp,
		PreRun: PreRun,
		RunE: func(c *cobra.Command, args []string) error {
			return golang.Push()
		},
	}

	// required flags
	cmd.Flags().StringVarP(&settings.Repo, "repo", "r", "", `e.g., --repo https://demo-go.pkg.wepack.net/project/repo/ or -r https://demo-go.pkg.wepack.net/project/repo/`)
	cmd.Flags().StringVarP(&settings.Module, "module", "m", "", `e.g., --module github.com/coding-wepack/wepack-cli@v0.0.1 or -m github.com/coding-wepack/wepack-cli@v0.0.1`)
	cmd.Flags().StringVarP(&settings.Username, "username", "u", "", "e.g., --username test or -u test")
	cmd.Flags().StringVarP(&settings.Password, "password", "p", "", "e.g., --password test_pwd or -u test_pwd")

	// Mark flags as required
	_ = cmd.MarkFlagRequired("module")
	_ = cmd.MarkFlagRequired("repo")
	return cmd
}
