package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validating the config",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runValidate(cmd.Context())
	},
}

func runValidate(_ context.Context) error {
	err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	err = cfg.Spec.Validate()
	if err != nil {
		return err
	}

	log.Print("OK")

	return nil
}
