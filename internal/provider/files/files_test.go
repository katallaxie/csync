package files

import (
	"testing"

	p "github.com/katallaxie/csync/pkg/provider"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		opts   []Opt
		expect *provider
	}{
		{
			name:   "empty",
			opts:   []Opt{},
			expect: &provider{},
		},
		{
			name: "with folder",
			opts: []Opt{
				WithFolder("foo"),
			},
			expect: &provider{
				folder: "foo",
			},
		},
	}

	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			f := New(test.opts...)

			assert.NotNil(t, f)
			assert.Implements(t, (*p.Provider)(nil), f)
			assert.Equal(t, test.expect, f)
		})
	}
}
