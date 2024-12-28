package spec_test

import (
	"testing"

	"github.com/katallaxie/csync/pkg/spec"

	"github.com/stretchr/testify/assert"
)

func TestAWS(t *testing.T) {
	tests := []struct {
		desc string
		in   spec.App
		out  spec.App
	}{
		{
			desc: "aws",
			in:   spec.AWS(),
			out: spec.App{
				Name: "aws",
				Files: spec.Files{
					"~/.aws",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.out, tc.in)
		})
	}
}
