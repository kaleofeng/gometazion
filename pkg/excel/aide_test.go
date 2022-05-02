package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAide_NewFile(t *testing.T) {
	ast := assert.New(t)

	aide := NewAide()
	aide.NewFile()

	sheet, err := aide.CreateSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aide.WriteRows(sheet, data)
	err = aide.Save("../../temp/excel.xlsx")
	ast.NoError(err)
}

func TestAide_OpenFile(t *testing.T) {
	ast := assert.New(t)

	aide := NewAide()
	err := aide.OpenFile("../../temp/excel.xlsx")
	ast.NoError(err)

	sheet, err := aide.GetSheet("Info")
	ast.NoError(err)

	data := [][]string{
		{"foo", "18"},
		{"bar", "20"},
	}
	aide.WriteRows(sheet, data)
	err = aide.Save("../../temp/excel.xlsx")
	ast.NoError(err)
}
