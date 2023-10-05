package files

import (
	"testing"

	p "github.com/katallaxie/csync/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	f := New()

	assert.NotNil(t, f)
	assert.Implements(t, (*p.Backup)(nil), f)
}
