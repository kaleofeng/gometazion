package mz

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
func (excelKit *ExcelKit) Open() {
	fmt.Println("Excel Kit - Open")
	excelKit.file = xlsx.NewFile()
}

// Save save as a file
func (excelKit *ExcelKit) Save(filename string) {
	excelKit.file.Save(filename)
	fmt.Printf("Excel Kit - Save: filename(%v)\n", filename)
}

// Write write data into a sheet
func (excelKit *ExcelKit) Write(sheetName string, data interface{}) {
	fmt.Printf("Excel Kit - Write: sheetName(%v)\n", sheetName)

	sheet, err := excelKit.file.AddSheet(sheetName)
	if err != nil {
		fmt.Printf("Excel Kit - Write: add sheet failed, err(%v)\n", err)
		return
	}

	value := reflect.ValueOf(data)
	switch value.Kind() {
	case reflect.Interface:
		excelKit.writeInterface(sheet, value)
	case reflect.String:
		excelKit.writeString(sheet, value)
	case reflect.Struct:
		excelKit.writeStruct(sheet, value)
	case reflect.Slice:
		excelKit.writeSlice(sheet, value)
	}

	fmt.Printf("Excel Kit - Write: success, sheetName(%v)\n", sheetName)
}

func (excelKit *ExcelKit) writeInterface(sheet *xlsx.Sheet, value reflect.Value) {
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (excelKit *ExcelKit) writeString(sheet *xlsx.Sheet, value reflect.Value) {
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = fmt.Sprintf("%v", value.Interface())
}

func (excelKit *ExcelKit) writeStruct(sheet *xlsx.Sheet, value reflect.Value) {
	row := sheet.AddRow()

	//valueType := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		//fieldType := valueType.Field(i)
		filedValue := fmt.Sprintf("%v", field.Interface())

		cell := row.AddCell()
		cell.Value = filedValue
	}
}

func (excelKit *ExcelKit) writeSlice(sheet *xlsx.Sheet, value reflect.Value) {
	for i := 0; i < value.Len(); i++ {
		each := value.Index(i)
		switch each.Kind() {
		case reflect.Interface:
			excelKit.writeInterface(sheet, each)
		case reflect.String:
			excelKit.writeString(sheet, each)
		case reflect.Struct:
			excelKit.writeStruct(sheet, each)
		}
	}
}
