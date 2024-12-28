package cmd

import (
	"context"
	"log"
	"os"

	"github.com/katallaxie/csync/pkg/spec"

	"github.com/katallaxie/pkg/utils/files"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	log.Printf("initializing config (%s)", cfg.File)

	if err := spec.Write(spec.Default(), cfg.File, cfg.Flags.Force); err != nil {
		return err
	}

	log.Printf("creating config folder (%s)", cfg.Path)

	err := files.MkdirAll(cfg.Path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
