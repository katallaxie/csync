package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/katallaxie/csync/internal/checker"
	"github.com/katallaxie/csync/internal/config"
	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/internal/ui"
	"github.com/katallaxie/csync/pkg/homedir"
	"github.com/katallaxie/csync/pkg/plugin"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var cfg = config.New()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(RestoreCmd)
	RootCmd.AddCommand(UnlinkCmd)
	RootCmd.AddCommand(ValidateCmd)
	RootCmd.AddCommand(BackupCmd)
	RootCmd.AddCommand(AppsCmd)

	RootCmd.PersistentFlags().StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Root, "root", "r", cfg.Flags.Root, "run as root")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
	RootCmd.PersistentFlags().StringVarP(&cfg.Flags.Plugin, "plugin", "p", cfg.Flags.Plugin, "plugin")
	RootCmd.PersistentFlags().StringSliceVar(&cfg.Flags.Vars, "var", cfg.Flags.Vars, "variables")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
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
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
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

	var p provider.Provider

	f, err := cfg.Spec.Provider.GetFolder()
	if err != nil {
		return err
	}

	// configuring the default file provider as fallback
	p = files.New(files.WithFolder(f), files.WithHomeDir(homedir.Get()))

	opts := provider.Opts{
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

	// see https://github.com/charmbracelet/lipgloss/issues/73
	lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

	model := ui.NewModel(apps, p.Backup, opts)

	proc := tea.NewProgram(
		model,
		// tea.WithAltScreen(),
		tea.WithReportFocus(),
		tea.WithContext(ctx),
	)

	if _, err := proc.Run(); err != nil {
		return err
	}

	return nil
}
