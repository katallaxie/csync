package spec_test

import (
	"fmt"
	"testing"

	"github.com/katallaxie/csync/internal/spec"

	"github.com/stretchr/testify/assert"
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
			assert.NoError(t, err)
			assert.Equal(t, tc.out, s)
		})
	}
}

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
