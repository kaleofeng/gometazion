package zlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	src := []byte("abcdefghijklmnopqrstuvwxyz")
	dst, err := Compress(src)
	ast.NoError(err)
	ast.True(len(dst) > 0)

	ori, err := Uncompress(dst)
	ast.NoError(err)
	ast.Equal(ori, src)
}
