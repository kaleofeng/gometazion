package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitState_Set(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		bit1 int
		bit2 int
		val  uint64
	}{
		{0, 64, 1 << 0},
		{63, 64, 1 << 63},
		{64, 65, 0},
		{0, 0, 1 << 0},
		{63, 63, 1 << 63},
		{0, 1, 1<<0 | 1<<1},
		{1, 63, 1<<1 | 1<<63},
	}

	for _, tc := range tcs {
		var bs BitState = 0
		bs.Set(tc.bit1).Set(tc.bit2)
		ast.Equal(tc.val, bs.Value())
	}
}

func TestBitState_Clear(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		bit1 int
		bit2 int
		val  uint64
	}{
		{0, 64, 1 << 0},
		{63, 64, 1 << 63},
		{64, 65, 0},
		{0, 0, 1 << 0},
		{63, 63, 1 << 63},
		{0, 1, 1<<0 | 1<<1},
		{1, 63, 1<<1 | 1<<63},
	}

	for _, tc := range tcs {
		var bs = BitState(tc.val)
		bs.Clear(tc.bit1).Clear(tc.bit2)
		ast.Equal(uint64(0), bs.Value())
	}
}

func TestBitState_Test(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	tcs := []struct {
		bit1 int
		bit2 int
		bit3 int
		bit4 int
		val  uint64
	}{
		{0, 0, 1, 63, 1 << 0},
		{63, 63, 62, 0, 1 << 63},
		{0, 1, 2, 63, 1<<0 | 1<<1},
		{1, 63, 0, 62, 1<<1 | 1<<63},
	}

	for _, tc := range tcs {
		var bs = BitState(tc.val)
		ast.True(bs.Test(tc.bit1))
		ast.True(bs.Test(tc.bit2))
		ast.False(bs.Test(tc.bit3))
		ast.False(bs.Test(tc.bit4))
	}
}
