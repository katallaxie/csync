package checker

import (
	"context"

	"github.com/katallaxie/csync/pkg/config"
)

type checker struct {
	opts *Opts
}

// Checker ...
type Checker interface {
	Check(context.Context, *config.Config) error
}

// Opt ...
type Opt func(*Opts)

// Func ...
type Func func(context.Context, *config.Config) error

// Opts ...
type Opts struct {
	funcs []Func
}

// Configure ...
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

	return &checker{
		opts: options,
	}
}

// Ready ...
func (c *checker) Check(ctx context.Context, cfg *config.Config) error {
	for _, fn := range c.opts.funcs {
		if err := fn(ctx, cfg); err != nil {
			return err
		}
	}

	return nil
}
