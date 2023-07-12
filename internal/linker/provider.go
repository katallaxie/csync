package linker

import (
	"fmt"

	"github.com/katallaxie/pkg/utils/files"
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
