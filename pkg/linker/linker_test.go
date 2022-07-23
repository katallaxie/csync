package linker_test

import (
	"context"
	"testing"

	"github.com/katallaxie/csync/pkg/linker"
	"github.com/katallaxie/csync/pkg/spec"

	"github.com/stretchr/testify/assert"
)

func TestBackup(t *testing.T) {
	var tests = []struct {
		desc        string
		app         *spec.App
		opts        []linker.Opt
		force       bool
		dry         bool
		expectedErr error
	}{
		{
			app: &spec.App{},
			opts: []linker.Opt{
				linker.WithProvider(&spec.Provider{}),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			l := linker.New(tc.opts...)
			err := l.Backup(ctx, tc.app, tc.force, tc.dry)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
