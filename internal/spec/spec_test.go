package spec_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/katallaxie/csync/internal/spec"

	"github.com/stretchr/testify/assert"
)

func Test_ProviderFilePath(t *testing.T) {
	tests := []struct {
		desc        string
		p           *spec.Provider
		f           string
		expected    string
		expectedErr error
	}{
		{
			p:           &spec.Provider{},
			f:           "foo.txt",
			expected:    "",
			expectedErr: fmt.Errorf("unknown provider"),
		},
		{
			p: &spec.Provider{
				Name: "file",
				Path: "/root",
			},
			f:           "foo.txt",
			expected:    "/root/csync/foo.txt",
			expectedErr: nil,
		},
		{
			p: &spec.Provider{
				Name:      "file",
				Path:      "/root",
				Directory: "/foo",
			},
			f:           "foo.txt",
			expected:    "/root/foo/foo.txt",
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			path, err := tc.p.GetFilePath(tc.f)
			assert.Equal(t, path, tc.expected)

			if tc.expectedErr != nil {
				assert.Error(t, err, tc.expectedErr)
			}
		})
	}
}

func Test_LoadSpec(t *testing.T) {
	tests := []struct {
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
    - "/Library/Preferences/"`,
			expected: &spec.Spec{
				Version: 1,
				Provider: spec.Provider{
					Name: "icloud",
				},
				Apps: []spec.App{
					{
						Name:  "nano",
						Files: []string{"/Library/Preferences/"},
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

			err = os.WriteFile(filepath.Join([]string{tempDir, "spec.yml"}...), content, 0o644)
			assert.NoError(t, err)

			s, err := spec.Load(filepath.Join([]string{tempDir, "spec.yml"}...))
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, s)
		})
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc        string
		spec        string
		expected    bool
		expectedErr error
	}{
		{
			desc: "valid",
			spec: `version: 1
path: /foo
`,
			expected:    true,
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := spec.Load(tc.spec)
			assert.Error(t, err, tc.expectedErr)
		})
	}

}
