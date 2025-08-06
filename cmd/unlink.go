package cmd

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/katallaxie/csync/internal/provider/files"
	"github.com/katallaxie/csync/internal/ui"
	pctx "github.com/katallaxie/csync/internal/ui/context"
	"github.com/katallaxie/csync/pkg/homedir"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var UnlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink the local files from the cloud",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
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
	model := ui.NewModel(pctx.WithContext(ctx), apps, p.Unlink, opts)

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
