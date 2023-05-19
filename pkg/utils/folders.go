package utils

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/andersnormal/pkg/utils/files"
)

// ICloudFolder ...
func ICloudFolder() (string, error) {
	path, err := files.ExpandHomeFolder("~/Library/Mobile Documents/com~apple~CloudDocs/")
	if err != nil {
		return "", err
	}

	ok, err := files.IsDir(path)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", fmt.Errorf("unable to find iCloud Drive")
	}

	return path, err
}

// DropboxFolder ...
func DropboxFodler() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	path := filepath.Join(usr.HomeDir, ".dropbox/host.db")

	file, err := os.OpenFile(filepath.Clean(path), os.O_RDWR, 0o600)
	if err != nil {
		panic(err)
	}

	defer func() { _ = file.Close() }()

	bb, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bb), "\n")

	dec, err := b64.URLEncoding.DecodeString(lines[1])
	if err != nil {
		return "", err
	}

	return string(dec), err
}
