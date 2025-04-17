package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/pkg/homedir"
	"github.com/katallaxie/csync/pkg/plugin"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore files from the cloud to the local machine",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
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

	err = cfg.Spec.Validate()
	if err != nil {
		return err
	}

	log.Printf("Restoring apps ...")

	if cfg.Flags.Dry {
		log.Printf("Running in Dry-Mode ...")
	}

	var p provider.Provider

	f, err := cfg.Spec.Provider.GetFolder()
	if err != nil {
		return err
	}

	// configuring the default file provider as fallback
	p = files.New(files.WithFolder(f), files.WithHomeDir(homedir.Get()))

	opts := &provider.Opts{
		Force: cfg.Flags.Force,
		Dry:   cfg.Flags.Dry,
		Root:  cfg.Flags.Root,
	}

	if cfg.Flags.Plugin != "" {
		m := plugin.Meta{Path: cfg.Flags.Plugin}
		f := m.Factory(ctx)

		p, err = f()
		if err != nil {
			return err
		}
	}

	defer p.Close()

	apps := cfg.Spec.GetApps(spec.List()...)
	for i := range apps {
		if err := p.Restore(ctx, &apps[i], opts); err != nil {
			return err
		}
	}

	return nil
}
