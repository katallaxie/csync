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
		expect *providerImpl
	}{
		{
			name:   "empty",
			opts:   []Opt{},
			expect: &providerImpl{},
		},
		{
			name: "with folder",
			opts: []Opt{
				WithFolder("foo"),
			},
			expect: &providerImpl{
				folder: "foo",
			},
		},
	}

	for _, tt := range tests {
		test := tt

		t.Run(test.name, func(t *testing.T) {
			f := New(test.opts...)

			assert.NotNil(t, f)
			assert.Implements(t, (*p.Provider)(nil), f)
			assert.Equal(t, test.expect, f)
		})
	}
}
