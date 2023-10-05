package icloud

import (
	"fmt"

	"github.com/katallaxie/pkg/utils/files"
)

// ErrNoICloudDrive is returned when the iCloud Drive folder cannot be found.
var ErrNoICloudDrive = fmt.Errorf("unable to find iCloud Drive")

type provider struct{}

// New ...
func New() *provider {
	return &provider{}
}

// Folder ...
func (p *provider) Folder() (string, error) {
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
