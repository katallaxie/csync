package files

import (
	p "github.com/katallaxie/csync/internal/provider"
	"github.com/katallaxie/pkg/utils/files"
)

type provider struct{}

var _ p.Backup = (*provider)(nil)

// New ...
func New() *provider {
	return &provider{}
}

// Folder ...
func (p *provider) Folder(f string) (string, error) {
	return files.ExpandHomeFolder(f)
}
