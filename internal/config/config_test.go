package config_test

import (
	"os"
	"testing"

	"github.com/katallaxie/csync/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := config.New()

	assert.NotNil(t, c)
	assert.Equal(t, c.File, ".csync.yml")
	assert.NotNil(t, c.Stderr)
	assert.NotNil(t, c.Stdout)
	assert.NotNil(t, c.Stdout)
	assert.Equal(t, c.Stderr, os.Stderr)
	assert.Equal(t, c.Stdout, os.Stdout)
	assert.Equal(t, c.Stderr, os.Stderr)
	assert.NotNil(t, c.Spec)

	h, err := c.HomeDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, h)

	cwd, err := c.Cwd()
	assert.NoError(t, err)
	assert.NotEmpty(t, cwd)
}

func TestUsePlugin(t *testing.T) {
	t.Parallel()

	c := config.New()
	c.Plugin = "dummy"

	assert.NotNil(t, c)
	assert.Equal(t, c.Plugin, "dummy")
	assert.True(t, c.UsePlugin())
}

func TestVars(t *testing.T) {
	t.Parallel()

	c := config.New()
	c.Flags.Vars = []string{"foo=bar"}

	assert.NotNil(t, c)
	assert.Equal(t, c.Flags.Vars, []string{"foo=bar"})
}
