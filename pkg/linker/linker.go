package linker

import (
	"context"

	"github.com/katallaxie/csync/pkg/spec"
)

type linker struct {
	opts *Opts
}

// Linker ...
type Linker interface {
	Link(context.Context, *spec.App) error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct{}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// New ...
func New(opts ...Opt) Linker {
	options := new(Opts)
	options.Configure(opts...)

	return &linker{
		opts: options,
	}
}

// Link ...
func (l *linker) Link(ctx context.Context, app *spec.App) error {
	return nil
}
