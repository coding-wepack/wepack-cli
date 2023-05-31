package main

import (
	"github.com/coding-wepack/wepack-cli/cmd/require"
	"github.com/spf13/cobra"

	"github.com/coding-wepack/wepack-cli/pkg/settings"
)

const globalUsage = `The WePack Artifacts Manager Client

The migrate argument must be an artifact type, available now:
- go

Common actions for wp:
- wp go: Manager go artifacts
`

func newRootCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "wp [TYPE]",
		Short:        "The WePack Artifacts Manager Client.",
		Long:         globalUsage,
		Args:         require.MinimumNArgs(1),
		SilenceUsage: true,
	}

	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.PersistentFlags().BoolVarP(&settings.Verbose, "verbose", "v", false, "Make the operation more talkative")

	cmd.AddCommand(
		newVersionCmd(),
		newGoCmd(),
	)

	return cmd, nil
}
