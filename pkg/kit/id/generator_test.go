package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	ig := NewGenerator(1000)
	ast.Equal(int64(1000), ig.CurrentId())
	ast.Equal(int64(1001), ig.NextId())
	ast.Equal(int64(1001), ig.CurrentId())
}
