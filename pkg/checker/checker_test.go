package checker_test

import (
	"context"
	"testing"

	"github.com/katallaxie/csync/pkg/checker"
	"github.com/katallaxie/csync/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := checker.New()
	assert.NotNil(t, c)
	assert.NoError(t, c.Check(context.Background(), nil))
}

func TestCheck(t *testing.T) {
	var tests = []struct {
		desc string
		fn   checker.Func
		cfg  *config.Config
	}{
		{
			desc: "no checks",
			fn:   func(ctx context.Context, c *config.Config) error { return nil },
			cfg:  nil,
		},
	}

	for _, tt := range tests {
		test := tt

		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			c := checker.New(checker.WithChecks(test.fn))
			assert.NotNil(t, c)
			assert.NoError(t, c.Check(context.Background(), test.cfg))
		})
	}
}
