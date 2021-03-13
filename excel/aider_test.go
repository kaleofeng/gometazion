package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_NewFile(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	aider := Aider{}
	aider.NewFile()

	sheet, err := aider.CreateSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aider.WriteRows(sheet, data)
	err = aider.Save("test.xlsx")
	ast.NoError(err)
}

func TestHelper_OpenFile(t *testing.T) {
	t.Parallel()
	ast := assert.New(t)

	aider := Aider{}
	err := aider.OpenFile("test.xlsx")
	ast.NoError(err)

	sheet, err := aider.GetSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aider.WriteRows(sheet, data)
	err = aider.Save("test.xlsx")
	ast.NoError(err)
}
