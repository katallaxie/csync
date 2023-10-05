package icloud

import (
	"testing"

	p "github.com/katallaxie/csync/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	i := New()

	assert.NotNil(t, i)
	assert.Implements(t, (*p.Backup)(nil), i)
}
