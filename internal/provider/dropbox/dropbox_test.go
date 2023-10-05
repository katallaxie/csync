package dropbox

import (
	"testing"

	p "github.com/katallaxie/csync/internal/provider"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	d := New()

	assert.NotNil(t, d)
	assert.Implements(t, (*p.Backup)(nil), d)
}
