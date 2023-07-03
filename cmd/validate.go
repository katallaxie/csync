package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:  "validate",
	Long: "validate",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate(cmd.Context())
	},
}

func runValidate(ctx context.Context) error {
	return nil
}
