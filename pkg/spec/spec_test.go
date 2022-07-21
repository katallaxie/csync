package spec_test

import (
	"testing"

	"github.com/katallaxie/csync/pkg/spec"
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
