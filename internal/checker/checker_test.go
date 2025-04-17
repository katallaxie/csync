package checker_test

import (
	"context"
	"testing"

	"github.com/katallaxie/csync/internal/checker"
	"github.com/katallaxie/csync/internal/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := checker.New()
	assert.NotNil(t, c)
	require.NoError(t, c.Check(t.Context(), nil))
}

func TestCheck(t *testing.T) {
	tests := []struct {
		desc string
		fn   checker.Func
		cfg  *config.Config
	}{
		{
			desc: "no checks",
			fn:   func(_ context.Context, _ *config.Config) error { return nil },
			cfg:  nil,
		},
	}

	for _, tt := range tests {
		test := tt

		t.Run(test.desc, func(t *testing.T) {
			c := checker.New(checker.WithChecks(test.fn))
			assert.NotNil(t, c)
			require.NoError(t, c.Check(t.Context(), test.cfg))
		})
	}
}
