package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/internal/linker"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files from the cloud to the local machine",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRestore(cmd.Context())
	},
}

func runRestore(ctx context.Context) error {
	err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	cfg.Lock()
	defer cfg.Unlock()

	if cfg.Flags.Verbose {
		log.Print("Restore files...")
	}

	opts := []linker.Opt{linker.WithProvider(cfg.Spec.Provider)}
	if cfg.Flags.Verbose {
		opts = append(opts, linker.WithVerbose())
	}

	l := linker.New(opts...)

	for _, a := range cfg.Spec.Apps {
		a := a
		if cfg.Flags.Verbose {
			log.Printf("Restoring %s", a.Name)
		}

		if err := l.Restore(ctx, &a, cfg.Flags.Force, cfg.Flags.Dry); err != nil {
			log.Panic(err)
		}
	}

	return nil
}
