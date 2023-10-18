package provider

import "github.com/katallaxie/csync/internal/spec"

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
	Backup(app *spec.App, opts *Opts) error
	// Restore an app.
	Restore(app *spec.App, opts *Opts) error
	// Link an app.
	Link(app *spec.App, opts *Opts) error
	// Unlink an app.
	Unlink(app *spec.App, opts *Opts) error
	// Close is a function to call before finalizing any action.
	// It returns an error if the process fails to gracefully finish.
	Close() error
}

// Unimplemented is the default implementation.
type Unimplemented struct{}

// Backup ...
func (p *Unimplemented) Backup(app *spec.App, opts *Opts) error {
	return nil
}

// Restore ...
func (p *Unimplemented) Restore(app *spec.App, opts *Opts) error {
	return nil
}

// Unlink ...
func (p *Unimplemented) Unlink(app *spec.App, opts *Opts) error {
	return nil
}

// Link ...
func (p *Unimplemented) Link(app *spec.App, opts *Opts) error {
	return nil
}
