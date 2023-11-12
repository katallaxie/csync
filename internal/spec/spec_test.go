package spec_test

import (
	"fmt"
	"testing"

	"github.com/katallaxie/csync/internal/spec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		desc string
		in   []byte
		out  *spec.Spec
		err  error
	}{
		{
			desc: "valid",
			in:   []byte(`version: 1`),
			out:  &spec.Spec{Version: 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			s := spec.Default()
			err := s.UnmarshalYAML(tc.in)
			require.NoError(t, err)
			assert.Equal(t, tc.out, s)
		})
	}
}

func Test_GetApps(t *testing.T) {
	tests := []struct {
		desc     string
		in       []byte
		out      []spec.App
		defaults []spec.App
	}{
		{
			desc: "default apps",
			in: []byte(
				`version: 1
apps:
  - 
    name: "example"
    files:
      - "/workspaces/csync/examples/example.txt"
`),
			out: append([]spec.App{
				{
					Name: "example",
					Files: []string{
						"/workspaces/csync/examples/example.txt",
					},
				},
			}, spec.List()...),
			defaults: spec.List(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			s := spec.Default()

			err := s.UnmarshalYAML(tc.in)
			require.NoError(t, err)

			assert.Equal(t, s.GetApps(tc.defaults...), tc.out)
		})
	}
}

func Test_ProviderFolder(t *testing.T) {
	tests := []struct {
		desc        string
		p           *spec.Provider
		expected    string
		expectedErr error
	}{
		{
			p:           &spec.Provider{},
			expected:    "",
			expectedErr: fmt.Errorf("unknown provider"),
		},
		{
			p: &spec.Provider{
				Name: "file",
				Path: "/root",
			},
			expected:    "/root/csync",
			expectedErr: nil,
		},
		{
			p: &spec.Provider{
				Name:      "file",
				Path:      "/root",
				Directory: "/foo",
			},
			expected:    "/root/foo",
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			path, err := tc.p.GetFolder()
			assert.Equal(t, tc.expected, path)

			if tc.expectedErr != nil {
				require.ErrorContains(t, err, tc.expectedErr.Error())
			}
		})
	}
}
