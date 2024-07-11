package checker

import (
	"context"

	"github.com/katallaxie/csync/internal/config"
)

var _ Checker = (*checkerImpl)(nil)

type checkerImpl struct {
	opts *Opts
}

// Checker is the interface for the checker to implement.
// This checks the configuration for correctness.
type Checker interface {
	Check(context.Context, *config.Config) error
}

// Opt is a functional option for the checker
type Opt func(*Opts)

// Func is a function to check the configuration
type Func func(context.Context, *config.Config) error

// Opts are the options for the checker
type Opts struct {
	funcs []Func
}

// Configure is a method to apply multiple options
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithChecks ...
func WithChecks(funcs ...Func) Opt {
	return func(o *Opts) {
		o.funcs = append(o.funcs, funcs...)
	}
}

// New ...
func New(opts ...Opt) Checker {
	options := new(Opts)
	options.Configure(opts...)

	return &checkerImpl{
		opts: options,
	}
}

// Ready ...
func (c *checkerImpl) Check(ctx context.Context, cfg *config.Config) error {
	for _, fn := range c.opts.funcs {
		if err := fn(ctx, cfg); err != nil {
			return err
		}
	}

	return nil
}
