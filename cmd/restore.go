package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:  "restore",
	Long: "restore",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRestore(cmd.Context())
	},
}

func runRestore(ctx context.Context) error {
	return nil
}
