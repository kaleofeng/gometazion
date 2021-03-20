package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIdGenerator(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	ig := NewIdGenerator(1000)
	ast.Equal(ig.CurrentId(), int32(1000))
	ast.Equal(ig.NextId(), int32(1001))
	ast.Equal(ig.CurrentId(), int32(1001))
}
