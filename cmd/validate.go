package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/pkg/spec"
	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:  "validate",
	Long: "validate",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate(cmd.Context())
	},
}

func runValidate(_ context.Context) error {
	s, err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	err = spec.Validate(s)
	if err != nil {
		return err
	}

	if cfg.Flags.Verbose {
		log.Print("OK")
	}

	return nil
}
