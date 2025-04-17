package provider

import (
	"context"

	"github.com/katallaxie/csync/pkg/spec"
)

// Opts are the options.
type Opts struct {
	// Dry toggles the dry run mode.
	Dry bool
	// Force toggles the force mode.
	Force bool
	// Root runs the command as root.
	Root bool
}

// Provider is the provider interface.
type Provider interface {
	// Backup an app.
	Backup(ctx context.Context, app *spec.App, opts *Opts) error
	// Restore an app.
	Restore(ctx context.Context, app *spec.App, opts *Opts) error
	// Link an app.
	Link(ctx context.Context, app *spec.App, opts *Opts) error
	// Unlink an app.
	Unlink(ctx context.Context, app *spec.App, opts *Opts) error
	// Close is a function to call before finalizing any action.
	// It returns an error if the process fails to gracefully finish.
	Close() error
}

var _ Provider = (*Unimplemented)(nil)

// Unimplemented is the default implementation.
type Unimplemented struct{}

// Backup is the backup function.
func (p *Unimplemented) Backup(_ context.Context, _ *spec.App, _ *Opts) error {
	return nil
}

// Restore is the restore function.
func (p *Unimplemented) Restore(_ context.Context, _ *spec.App, _ *Opts) error {
	return nil
}

// Unlink is the unlink function.
func (p *Unimplemented) Unlink(_ context.Context, _ *spec.App, _ *Opts) error {
	return nil
}

// Link is the link function.
func (p *Unimplemented) Link(_ context.Context, _ *spec.App, _ *Opts) error {
	return nil
}

// Close implements the closer interface.
func (p *Unimplemented) Close() error {
	return nil
}
