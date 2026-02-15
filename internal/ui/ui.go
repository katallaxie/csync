package ui

import (
	"context"
	"fmt"
	"strings"

	pctx "github.com/katallaxie/csync/internal/ui/context"
	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"
	"github.com/katallaxie/pkg/utilx"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Cmd is the command to run the UI.
type Cmd func(ctx context.Context, app spec.App, opts provider.Opts) error

// Mode le is the model for the UI.
type Model struct {
	apps     []spec.App
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
	cmd      Cmd
	opts     provider.Opts
	ctx      *pctx.ProgramContext
}

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle           = lipgloss.NewStyle().Margin(1, 2)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorCross          = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).SetString("✗")
)

// NewModel creates a new Model.
func NewModel(ctx *pctx.ProgramContext, apps []spec.App, cmd Cmd, opts provider.Opts) Model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	return Model{
		apps:     apps,
		cmd:      cmd,
		spinner:  s,
		progress: p,
		opts:     opts,
		ctx:      ctx,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(installApp(m.ctx.Context(), m.apps[m.index], m.cmd, m.opts), m.spinner.Tick)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	case pctx.InstalledPkgMessage:
		if m.index >= len(m.apps)-1 {
			m.done = true
			return m, tea.Sequence(
				utilx.IfElse(
					utilx.Empty(msg.Err),
					tea.Printf("%s %s", checkMark, msg.App.Name),
					tea.Printf("%s %s", errorCross, msg.App.Name),
				),
				tea.Quit,
			)
		}

		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.apps)))
		msgCmd := utilx.IfElse(
			utilx.Empty(msg.Err),
			tea.Printf("%s %s", checkMark, msg.App.Name),
			tea.Printf("%s %s", errorCross, msg.App.Name),
		)

		cmds = append(cmds, progressCmd, msgCmd, installApp(m.ctx.Context(), m.apps[m.index], m.cmd, m.opts))
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	n := len(m.apps)
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render("Done!")
	}

	pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+pkgCount))

	pkgName := currentPkgNameStyle.Render(m.apps[m.index].Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render(pkgName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+pkgCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + pkgCount
}

func installApp(ctx context.Context, app spec.App, cmd Cmd, opts provider.Opts) tea.Cmd {
	return func() tea.Msg {
		return pctx.InstalledPkgMessage{
			App: app,
			Err: cmd(ctx, app, opts),
		}
	}
}
