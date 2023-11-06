package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/internal/spec"
	"github.com/katallaxie/csync/pkg/plugin"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/spf13/cobra"
)

var UnlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink the local files from the cloud",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUnlink(cmd.Context())
	},
}

func runUnlink(ctx context.Context) error {
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

	log.Printf("Unlinking apps ...")

	if cfg.Flags.Dry {
		log.Printf("Running in Dry-Mode ...")
	}

	var p provider.Provider

	f, err := cfg.Spec.Provider.GetFolder()
	if err != nil {
		return err
	}

	// configuring the default file provider as fallback
	p = files.New(files.WithFolder(f))

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
		log.Printf("Unlink '%s'", apps[i].Name)

		if err := p.Unlink(&apps[i], opts); err != nil {
			return err
		}
	}

	return nil
}
