package cmd

import (
	"context"
	"log"
	"os"

	"github.com/andersnormal/pkg/utils/files"
	"github.com/katallaxie/csync/pkg/spec"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	if cfg.Flags.Verbose {
		log.Printf("initializing config (%s)", cfg.File)
	}

	if err := spec.Write(spec.Default(), cfg.File, cfg.Flags.Force); err != nil {
		return err
	}

	if cfg.Flags.Verbose {
		log.Printf("creating config folder (%s)", cfg.Path)
	}

	err := files.MkdirAll(cfg.Path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
