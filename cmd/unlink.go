package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var UnlinkCmd = &cobra.Command{
	Use:  "unlink",
	Long: "unlink",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUnlink(cmd.Context())
	},
}

func runUnlink(ctx context.Context) error {
	return nil
}
