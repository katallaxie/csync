package plugins

import (
	"context"
	"errors"
	"plugin"

	"github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"
)

var Unknown = "unknown"

var ErrUnimplemented = errors.New("not implemented")

var _ Plugin = (*UnimplementedPlugin)(nil)

type UnimplementedPlugin struct{}

// Name returns the name of the plugin.
func (u *UnimplementedPlugin) Name() string {
	return Unknown
}

// Description returns a brief description of the plugin.
func (u *UnimplementedPlugin) Description() string {
	return Unknown
}

// Version returns the version of the plugin.
func (u *UnimplementedPlugin) Version() string {
	return Unknown
}

// Backup performs a backup of the application data.
func (u *UnimplementedPlugin) Backup(_ context.Context, _ spec.App, _ provider.Opts) error {
	return ErrUnimplemented
}

// Restore restores the application data from a backup.
func (u *UnimplementedPlugin) Restore(_ context.Context, _ spec.App, _ provider.Opts) error {
	return ErrUnimplemented
}

// Link applications with the plugin.
func (u *UnimplementedPlugin) Link(_ context.Context, _ spec.App, _ provider.Opts) error {
	return ErrUnimplemented
}

// Unlink removes the link with the plugin.
func (u *UnimplementedPlugin) Unlink(_ context.Context, _ spec.App, _ provider.Opts) error {
	return ErrUnimplemented
}

// Plugin interface defines the methods that a plugin must implement.
type Plugin interface {
	// Name returns the name of the plugin.
	Name() string
	// Description returns a brief description of the plugin.
	Description() string
	// Version returns the version of the plugin.
	Version() string
	// Backup performs a backup of the application data.
	Backup(ctx context.Context, app spec.App, opts provider.Opts) error
	// Restore restores the application data from a backup.
	Restore(ctx context.Context, app spec.App, opts provider.Opts) error
	// Link applications with the plugin.
	Link(ctx context.Context, app spec.App, opts provider.Opts) error
	// Unlink removes the link with the plugin.
	Unlink(ctx context.Context, app spec.App, opts provider.Opts) error
}

// Load loads the plugin with the given name and returns it.
func Load(path string) (Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	p, ok := symPlugin.(Plugin)
	if !ok {
		return nil, err
	}

	return p, nil
}
