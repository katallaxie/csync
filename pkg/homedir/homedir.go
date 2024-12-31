package homedir

import (
	"bytes"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// Get returns the home directory of the current user with the help of
// environment variables depending on the target operating system.
// Returned path should be used with "path/filepath" to form new paths.
//
// On non-Windows platforms, it falls back to nss lookups, if the home
// directory cannot be obtained from environment-variables.
//
// If linking statically with cgo enabled against glibc, ensure the
// osusergo build tag is used.
//
// If needing to do nss lookups, do not disable cgo or set osusergo.
func Get() string {
	home, _ := os.UserHomeDir()

	if home == "" && runtime.GOOS != "windows" {
		if u, err := user.Current(); err == nil {
			return u.HomeDir
		}
	}

	return home
}

// Expand is a helper function to expand the home folder of a path.
func Expand(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	var buffer bytes.Buffer
	_, err := buffer.WriteString(Get())
	if err != nil {
		return "", err
	}

	_, err = buffer.WriteString(strings.TrimPrefix(path, "~"))
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
