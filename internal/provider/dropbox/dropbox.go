package dropbox

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ErrNoDropbox is returned when the Dropbox folder cannot be found.
var ErrNoDropbox = fmt.Errorf("unable to find Dropbox")

type provider struct{}

// New ...
func New() *provider {
	return &provider{}
}

// Folder ...
func (p *provider) Folder() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	path := filepath.Join(usr.HomeDir, ".dropbox/host.db")

	file, err := os.OpenFile(filepath.Clean(path), os.O_RDWR, 0o600)
	if err != nil {
		return "", err
	}

	defer func() { _ = file.Close() }()

	bb, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(bb), "\n")

	dec, err := b64.URLEncoding.DecodeString(lines[1])
	if err != nil {
		return "", err
	}

	return string(dec), err
}
