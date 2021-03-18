package excel

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx"
)

// Aider provides excel api wrapper.
type Aider struct {
	file *xlsx.File
}

// NewAider new an instance.
func NewAider() *Aider {
	return &Aider{}
}

// NewFile create a new file.
func (aider *Aider) NewFile() {
	aider.file = xlsx.NewFile()
}

// OpenFile open an existing file.
func (aider *Aider) OpenFile(filePath string) (err error) {
	aider.file, err = xlsx.OpenFile(filePath)
	return
}

// Save save as a file.
func (aider *Aider) Save(filePath string) (err error) {
	err = aider.file.Save(filePath)
	return
}

// CreateSheet create a new sheet.
func (aider *Aider) CreateSheet(sheetName string) (sheet *xlsx.Sheet, err error) {
	sheet, err = aider.file.AddSheet(sheetName)
	return
}

// GetSheet get an existing sheet.
func (aider *Aider) GetSheet(sheetName string) (sheet *xlsx.Sheet, err error) {
	sheet, ok := aider.file.Sheet[sheetName]
	if !ok {
		err = fmt.Errorf("no such sheet[%s]", sheetName)
	}
	return
}

// WriteRows write data into rows.
func (aider *Aider) WriteRows(sheet *xlsx.Sheet, data interface{}) {
	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		row := sheet.AddRow()
		aider.writeData(row, each)
	}
}

// WriteRow write data into a row.
func (aider *Aider) WriteRow(sheet *xlsx.Sheet, data interface{}) {
	row := sheet.AddRow()
	value := reflect.ValueOf(data)
	aider.writeData(row, value)
}

func (aider *Aider) writeData(row *xlsx.Row, value reflect.Value) {
	switch value.Kind() {
	case reflect.Interface:
		aider.writeInterface(row, value)
	case reflect.String:
		aider.writeString(row, value)
	case reflect.Struct:
		aider.writeStruct(row, value)
	case reflect.Slice:
		aider.writeSlice(row, value)
	}
}

func (aider *Aider) writeInterface(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (aider *Aider) writeString(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (aider *Aider) writeStruct(row *xlsx.Row, value reflect.Value) {
	// valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		// fieldType := valueType.Field(i)
		filedValue := fmt.Sprintf("%v", field.Interface())

		cell := row.AddCell()
		cell.Value = filedValue
	}
}

func (aider *Aider) writeSlice(row *xlsx.Row, value reflect.Value) {
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		switch each.Kind() {
		case reflect.Interface:
			aider.writeInterface(row, each)
		case reflect.String:
			aider.writeString(row, each)
		case reflect.Struct:
			aider.writeStruct(row, each)
		}
	}
}
