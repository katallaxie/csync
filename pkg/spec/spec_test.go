package spec_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/katallaxie/csync/pkg/spec"

	"github.com/stretchr/testify/assert"
)

func Test_FilePathFromProvider(t *testing.T) {
	var tests = []struct {
		desc        string
		p           *spec.Provider
		f           string
		expected    string
		expectedErr error
	}{}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {

		})
	}
}

func Test_LoadSpec(t *testing.T) {
	var tests = []struct {
		desc        string
		spec        string
		expected    *spec.Spec
		expectedErr error
	}{
		{
			spec: `version: 1
provider:
  name: icloud
# path: override the destination
apps:
  - 
    name: "nano"
    files:
    - "/Libary/Preferences/"`,
			expected: &spec.Spec{
				Version: 1,
				Provider: spec.Provider{
					Name: "icloud",
				},
				Apps: []spec.App{
					{
						Name:  "nano",
						Files: []string{"/Libary/Preferences/"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			tempDir, err := os.MkdirTemp(os.TempDir(), "empty_test")
			assert.NoError(t, err)

			defer func() { _ = os.RemoveAll(tempDir) }()

			content := []byte(tc.spec)

			fmt.Println(content)

			err = os.WriteFile(filepath.Join([]string{tempDir, "spec.yml"}...), content, 0644)
			assert.NoError(t, err)

			s, err := spec.Load(filepath.Join([]string{tempDir, "spec.yml"}...))
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, s)
		})
	}
}
