package excel

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx"
)

// ExcelKit provides excel api wrapper.
type ExcelKit struct {
	file *xlsx.File
}

// Open open a file
func (excelKit *ExcelKit) New() {
	fmt.Println("Excel Kit - New")
	excelKit.file = xlsx.NewFile()
}

// Save save as a file
func (excelKit *ExcelKit) Save(filename string) {
	excelKit.file.Save(filename)
	fmt.Printf("Excel Kit - Save: filename(%s)\n", filename)
}

// Write create a new sheet
func (excelKit *ExcelKit) CreateSheet(sheetName string) (sheet *xlsx.Sheet, err error) {
	fmt.Printf("Excel Kit - Create Sheet: sheetName(%s)\n", sheetName)

	sheet, err = excelKit.file.AddSheet(sheetName)
	return
}

// Write write data into rows
func (excelKit *ExcelKit) WriteRows(sheet *xlsx.Sheet, data interface{}) {
	fmt.Printf("Excel Kit - Write Rows\n")

	value := reflect.ValueOf(data)
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		row := sheet.AddRow()
		excelKit.writeData(row, each)
	}
}

// Write write data into a row
func (excelKit *ExcelKit) WriteRow(sheet *xlsx.Sheet, data interface{}) {
	fmt.Printf("Excel Kit - Write Row\n")

	row := sheet.AddRow()
	value := reflect.ValueOf(data)
	excelKit.writeData(row, value)
}

func (excelKit *ExcelKit) writeData(row *xlsx.Row, value reflect.Value) {
	switch value.Kind() {
	case reflect.Interface:
		excelKit.writeInterface(row, value)
	case reflect.String:
		excelKit.writeString(row, value)
	case reflect.Struct:
		excelKit.writeStruct(row, value)
	case reflect.Slice:
		excelKit.writeSlice(row, value)
	}
}

func (excelKit *ExcelKit) writeInterface(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (excelKit *ExcelKit) writeString(row *xlsx.Row, value reflect.Value) {
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (excelKit *ExcelKit) writeStruct(row *xlsx.Row, value reflect.Value) {
	// valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		// fieldType := valueType.Field(i)
		filedValue := fmt.Sprintf("%v", field.Interface())

		cell := row.AddCell()
		cell.Value = filedValue
	}
}

func (excelKit *ExcelKit) writeSlice(row *xlsx.Row, value reflect.Value) {
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		switch each.Kind() {
		case reflect.Interface:
			excelKit.writeInterface(row, each)
		case reflect.String:
			excelKit.writeString(row, each)
		case reflect.Struct:
			excelKit.writeStruct(row, each)
		}
	}
}
