package homing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAider_LoadIni(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	aider := NewAider()
	err := aider.LoadFromIni("../temp/homing.ini")
	ast.NoError(err)

	err = aider.ReplaceTextFile("../temp/homing.yaml")
	ast.NoError(err)
}
