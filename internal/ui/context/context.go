package context

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type contextKey int

const programContextKey contextKey = 0 // __local_user_context__

// NewProgramContext creates a new ProgramContext with default values.
func NewProgramContext() *ProgramContext {
	return &ProgramContext{}
}

// WithContext creates a new ProgramContext with the given context.
func WithContext(ctx context.Context) *ProgramContext {
	return &ProgramContext{
		ctx: ctx,
	}
}

// ProgramContext is the context for the program.
type ProgramContext struct {
	ScreenHeight      int
	ScreenWidth       int
	MainContentWidth  int
	MainContentHeight int

	ctx context.Context
}

// Context returns the context for the program.
func (c *ProgramContext) Context() context.Context {
	if c.ctx == nil {
		c.ctx = context.Background()
	}

	return c.ctx
}

// SetContext sets the context for the program.
func (c *ProgramContext) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// State is the state of a task.
type State = int

const (
	TaskStart State = iota
	TaskFinished
	TaskError
)

// Task is a task to be executed.
type Task struct {
	Id           string
	StartText    string
	FinishedText string
	State        State
	Error        error
	StartTime    time.Time
	FinishedTime *time.Time
	StartTask    func(task Task) tea.Cmd
}
