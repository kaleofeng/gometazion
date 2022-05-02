package excel

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx"
)

// Aide provides excel api wrapper.
type Aide struct {
	file *xlsx.File
}

// NewAide new an instance.
func NewAide() *Aide {
	return &Aide{}
}

// NewFile create a new file.
func (aide *Aide) NewFile() {
	aide.file = xlsx.NewFile()
}

// OpenFile open an existing file.
func (aide *Aide) OpenFile(filePath string) (err error) {
	aide.file, err = xlsx.OpenFile(filePath)
	return
}

// Save save as a file.
func (aide *Aide) Save(filePath string) (err error) {
	err = aide.file.Save(filePath)
	return
}

// CreateSheet create a new sheet.
func (aide *Aide) CreateSheet(sheetName string) (sheet *xlsx.Sheet, err error) {
	sheet, err = aide.file.AddSheet(sheetName)
	return
}

// GetSheet get an existing sheet.
func (aide *Aide) GetSheet(sheetName string) (sheet *xlsx.Sheet, err error) {
	sheet, ok := aide.file.Sheet[sheetName]
	if !ok {
		err = fmt.Errorf("no such sheet[%s]", sheetName)
	}
	return
}

// WriteRows write data into rows.
func (aide *Aide) WriteRows(sheet *xlsx.Sheet, data interface{}) {
	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		row := sheet.AddRow()
		aide.writeData(row, each)
	}
}

// WriteRow write data into a row.
func (aide *Aide) WriteRow(sheet *xlsx.Sheet, data interface{}) {
	row := sheet.AddRow()
	value := reflect.ValueOf(data)
	aide.writeData(row, value)
}

func (aide *Aide) writeData(row *xlsx.Row, value reflect.Value) {
	switch value.Kind() {
	case reflect.Interface:
		aide.writeInterface(row, value)
	case reflect.String:
		aide.writeString(row, value)
	case reflect.Struct:
		aide.writeStruct(row, value)
	case reflect.Slice:
		aide.writeSlice(row, value)
	}
}

func (aide *Aide) writeInterface(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (aide *Aide) writeString(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (aide *Aide) writeStruct(row *xlsx.Row, value reflect.Value) {
	// valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		// fieldType := valueType.Field(i)
		filedValue := fmt.Sprintf("%v", field.Interface())

		cell := row.AddCell()
		cell.Value = filedValue
	}
}

func (aide *Aide) writeSlice(row *xlsx.Row, value reflect.Value) {
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		switch each.Kind() {
		case reflect.Interface:
			aide.writeInterface(row, each)
		case reflect.String:
			aide.writeString(row, each)
		case reflect.Struct:
			aide.writeStruct(row, each)
		}
	}
}
