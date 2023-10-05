package icloud

import (
	"fmt"

	p "github.com/katallaxie/csync/internal/provider"
	"github.com/katallaxie/pkg/utils/files"
)

// ErrNoICloudDrive is returned when the iCloud Drive folder cannot be found.
var ErrNoICloudDrive = fmt.Errorf("unable to find iCloud Drive")

type provider struct{}

var _ p.Backup = (*provider)(nil)

// New ...
func New() *provider {
	return &provider{}
}

// Folder ...
func (p *provider) Folder(f string) (string, error) {
	path, err := files.ExpandHomeFolder("~/Library/Mobile Documents/com~apple~CloudDocs/")
	if err != nil {
		return "", err
	}

	ok, err := files.IsDir(path)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrNoICloudDrive
	}

	return path, err
}
