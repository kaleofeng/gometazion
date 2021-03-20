package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceBuffer_Push(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	pb := NewPieceBuffer(16)

	size := pb.Push([]byte("Test"))
	ast.Equal(4, size)
	ast.Equal(4, pb.GetPullLength())
	ast.Equal(12, pb.GetPushLength())

	size = pb.Push([]byte("Push"))
	ast.Equal(4, size)
	ast.Equal(8, pb.GetPullLength())
	ast.Equal(8, pb.GetPushLength())

	size = pb.Push([]byte("MoreDataCome"))
	ast.Equal(8, size)
	ast.Equal(16, pb.GetPullLength())
	ast.Equal(0, pb.GetPushLength())
}

func TestPieceBuffer_Pull(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	pb := NewPieceBuffer(16)

	size := pb.Push([]byte("TestPullOk"))
	ast.Equal(10, size)
	ast.Equal(10, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	out := make([]byte, 4)

	size = pb.Pull(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(6, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	size = pb.Pull(out)
	ast.Equal(4, size)
	ast.Equal("Pull", string(out))
	ast.Equal(2, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	size = pb.Pull(out)
	ast.Equal(2, size)
	ast.Equal("Ok", string(out[:size]))
	ast.Equal(0, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())
}

func TestPieceBuffer_Peek(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	pb := NewPieceBuffer(16)

	size := pb.Push([]byte("TestDataOk"))
	ast.Equal(10, size)
	ast.Equal(10, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	out := make([]byte, 4)

	size = pb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(10, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	size = pb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Test", string(out))
	ast.Equal(10, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())
}

func TestPieceBuffer_Skip(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	pb := NewPieceBuffer(16)

	size := pb.Push([]byte("TestDataOk"))
	ast.Equal(10, size)
	ast.Equal(10, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	size = pb.Skip(4)
	ast.Equal(4, size)
	ast.Equal(6, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())

	out := make([]byte, 4)

	size = pb.Peek(out)
	ast.Equal(4, size)
	ast.Equal("Data", string(out))
	ast.Equal(6, pb.GetPullLength())
	ast.Equal(6, pb.GetPushLength())
}

func TestPieceBuffer_Compact(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	pb := NewPieceBuffer(16)

	size := pb.Push([]byte("testdata"))
	ast.Equal(8, size)

	out := make([]byte, 4)

	size = pb.Pull(out)
	ast.Equal(4, size)
	ast.Equal(4, pb.GetPullLength())
	ast.Equal(8, pb.GetPushLength())

	pb.Compact()
	ast.Equal(4, pb.GetPullLength())
	ast.Equal(12, pb.GetPushLength())
}
