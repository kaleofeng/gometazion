package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingBuffer_Push(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	rb := NewRingBuffer(10)

	size := rb.Push([]byte("Test"))
	ast.Equal(4, size)
	ast.Equal(4, rb.GetPullLength())
	ast.Equal(6, rb.GetPushLength())

	size = rb.Push([]byte("Push"))
	ast.Equal(4, size)
	ast.Equal(8, rb.GetPullLength())
	ast.Equal(2, rb.GetPushLength())

	size = rb.Push([]byte("More"))
	ast.Equal(2, size)
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())
}

func TestRingBuffer_Pull(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	rb := NewRingBuffer(10)

	size := rb.Push([]byte("TestPullOk"))
	ast.Equal(10, size)
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())

	out := make([]byte, 4)

	size = rb.Pull(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(6, rb.GetPullLength())
	ast.Equal(4, rb.GetPushLength())

	size = rb.Pull(out)
	ast.Equal(4, size)
	ast.Equal("Pull", string(out))
	ast.Equal(2, rb.GetPullLength())
	ast.Equal(8, rb.GetPushLength())

	size = rb.Pull(out)
	ast.Equal(2, size)
	ast.Equal("Ok", string(out[:size]))
	ast.Equal(0, rb.GetPullLength())
	ast.Equal(10, rb.GetPushLength())
}

func TestRingBuffer_Peek(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	rb := NewRingBuffer(10)

	size := rb.Push([]byte("TestDataOk"))
	ast.Equal(10, size)
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())

	out := make([]byte, 4)

	size = rb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())

	size = rb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())
}

func TestRingBuffer_Skip(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	rb := NewRingBuffer(10)

	size := rb.Push([]byte("TestDataOk"))
	ast.Equal(10, size)
	ast.Equal(10, rb.GetPullLength())
	ast.Equal(0, rb.GetPushLength())

	size = rb.Skip(4)
	ast.Equal(4, size)
	ast.Equal(6, rb.GetPullLength())
	ast.Equal(4, rb.GetPushLength())

	out := make([]byte, 4)

	size = rb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Data", string(out))
	ast.Equal(6, rb.GetPullLength())
	ast.Equal(4, rb.GetPushLength())
}
