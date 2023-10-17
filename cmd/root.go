package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/internal/checker"
	"github.com/katallaxie/csync/internal/config"
	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/pkg/plugin"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/spf13/cobra"
)

var cfg = config.New()

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(RestoreCmd)
	RootCmd.AddCommand(UnlinkCmd)
	RootCmd.AddCommand(ValidateCmd)
	RootCmd.AddCommand(BackupCmd)

	RootCmd.PersistentFlags().StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Root, "root", "r", cfg.Flags.Root, "run as root")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.Plugin, "plugin", "p", cfg.Flags.Plugin, "plugin")
	RootCmd.PersistentFlags().StringSliceVar(&cfg.Flags.Vars, "var", cfg.Flags.Vars, "variables")

	RootCmd.SilenceErrors = true
}

func initConfig() {
	err := cfg.InitDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "csync",
	Short: "csync",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func checkEnv(ctx context.Context) error {
	c := checker.New(
		checker.WithChecks(checker.UseableEnv),
		checker.WithChecks(checker.UseSetup),
	)

	if err := c.Check(ctx, cfg); err != nil {
		return err
	}

	return nil
}

func runRoot(ctx context.Context) error {
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

	if cfg.Flags.Verbose {
		log.Printf("Backup apps ...")
	}

	var p provider.Provider
	p = files.New()

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

	for _, app := range cfg.Spec.GetApps() {
		log.Printf("Backup '%s", app.Name)

		if err := p.Backup(&app, opts); err != nil {
			return err
		}
	}

	return nil
}
