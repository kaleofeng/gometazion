package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAider_NewFile(t *testing.T) {
	ast := assert.New(t)

	aider := NewAider()
	aider.NewFile()

	sheet, err := aider.CreateSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aider.WriteRows(sheet, data)
	err = aider.Save("../../temp/excel.xlsx")
	ast.NoError(err)
}

func TestAider_OpenFile(t *testing.T) {
	ast := assert.New(t)

	aider := NewAider()
	err := aider.OpenFile("../../temp/excel.xlsx")
	ast.NoError(err)

	sheet, err := aider.GetSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aider.WriteRows(sheet, data)
	err = aider.Save("../../temp/excel.xlsx")
	ast.NoError(err)
}
