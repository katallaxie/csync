package utils

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/katallaxie/pkg/filex"
)

var (
	// ErrNoICloudDrive is returned when the iCloud Drive folder cannot be found.
	ErrNoICloudDrive = fmt.Errorf("unable to find iCloud Drive")
	// ErrNoDropbox is returned when the Dropbox folder cannot be found.
	ErrNoDropbox = fmt.Errorf("unable to find Dropbox folder")
	// ErrNoOpenCloud is returned when the OpenCloud folder cannot be found.
	ErrNoOpenCloud = fmt.Errorf("unable to find OpenCloud folder")
)

// ICloudFolder is the path to the iCloud Drive folder.
func ICloudFolder() (string, error) {
	path, err := filex.ExpandHomeFolder("~/Library/Mobile Documents/com~apple~CloudDocs/")
	if err != nil {
		return "", err
	}

	ok, err := filex.IsDir(path)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrNoICloudDrive
	}

	return path, err
}

// DropboxFolder is the path to the Dropbox folder.
func DropboxFolder() (string, error) {
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

// OpenCloudFolder is the path to the OpenCloud folder.
func OpenCloudFolder() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	path := filepath.Join(usr.HomeDir, "OpenCloud")
	ok, err := filex.IsDir(path)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", ErrNoOpenCloud
	}

	return path, nil
}
