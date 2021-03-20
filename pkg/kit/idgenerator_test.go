package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIdGenerator(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	ig := NewIdGenerator(1000)
	ast.Equal(int32(1000), ig.CurrentId())
	ast.Equal(int32(1001), ig.NextId())
	ast.Equal(int32(1001), ig.CurrentId())
}
