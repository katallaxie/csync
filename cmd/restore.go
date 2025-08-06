package cmd

import (
	"context"

	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/internal/ui"
	pctx "github.com/katallaxie/csync/internal/ui/context"
	"github.com/katallaxie/csync/pkg/homedir"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
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

	// see https://github.com/charmbracelet/lipgloss/issues/73
	lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

	apps := cfg.Spec.GetApps(spec.List()...)
	model := ui.NewModel(pctx.WithContext(ctx), apps, p.Restore, opts)

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
