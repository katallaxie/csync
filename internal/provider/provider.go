package provider

import "github.com/katallaxie/csync/internal/spec"

// Backup is the backup interface.
type Backup interface {
	// Folder returns the backup folder.
	Folder(f string) (string, error)
}

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
	// Backup a file.
	Backup(app *spec.App, opts *Opts) error
	// Restore a file.
	Restore(app *spec.App, opts *Opts) error
}
