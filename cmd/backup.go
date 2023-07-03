package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runBackup(cmd.Context())
	},
}

func runBackup(ctx context.Context) error {
	return nil
}
