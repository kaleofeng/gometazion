package homing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAide_LoadIni(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	aide := NewAide()
	err := aide.LoadFromIni("../../temp/homing.ini")
	ast.NoError(err)

	err = aide.ReplaceTextFile("../../temp/homing.yaml")
	ast.NoError(err)
}
