package files

import (
	"github.com/katallaxie/csync/internal/spec"
	p "github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/pkg/utils/files"
)

type provider struct {
	folder string

	p.Provider
}

var _ p.Provider = (*provider)(nil)

// Opt is the functional option for the provider.
type Opt func(*provider)

// WithFolder is configuring a specific folder for the provider.
func WithFolder(folder string) Opt {
	return func(p *provider) {
		p.folder = folder
	}
}

// Configure is configuring a set of options of the provider.
func (p *provider) Configure(opts ...Opt) {
	for _, o := range opts {
		o(p)
	}
}

// New ...
func New(opts ...Opt) *provider {
	return &provider{}
}

// Folder ...
func (p *provider) Folder(f string) (string, error) {
	return files.ExpandHomeFolder(f)
}

// Backup a file.
func (p *provider) Backup(app *spec.App, opts *p.Opts) error {
	return nil
}

// Restore a file.
func (p *provider) Restore(app *spec.App, opts *p.Opts) error {
	return nil
}
