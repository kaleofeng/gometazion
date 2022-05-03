package transform

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func DecodeValueFromBytes(data []byte, order binary.ByteOrder, value interface{}) error {
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, order, value); err != nil {
		return err
	}
	return nil
}

func EncodeValueToBytes(value interface{}, order binary.ByteOrder) ([]byte, error) {
	writer := bytes.NewBuffer([]byte{})
	if err := binary.Write(writer, order, value); err != nil {
		return []byte{}, err
	}
	return writer.Bytes(), nil
}

func BytesToInt8(data []byte, order binary.ByteOrder) (int8, error) {
	value := int8(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToUint8(data []byte, order binary.ByteOrder) (uint8, error) {
	value := uint8(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToInt16(data []byte, order binary.ByteOrder) (int16, error) {
	value := int16(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToUint16(data []byte, order binary.ByteOrder) (uint16, error) {
	value := uint16(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToInt32(data []byte, order binary.ByteOrder) (int32, error) {
	value := int32(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToUint32(data []byte, order binary.ByteOrder) (uint32, error) {
	value := uint32(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToInt64(data []byte, order binary.ByteOrder) (int64, error) {
	value := int64(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToUint64(data []byte, order binary.ByteOrder) (uint64, error) {
	value := uint64(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToFloat32(data []byte, order binary.ByteOrder) (float32, error) {
	value := float32(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func BytesToFloat64(data []byte, order binary.ByteOrder) (float64, error) {
	value := float64(0)
	err := DecodeValueFromBytes(data, order, &value)
	return value, err
}

func StringToInt8(data string) (int8, error) {
	value, err := strconv.ParseInt(data, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(value), nil
}

func StringToUint8(data string) (uint8, error) {
	value, err := strconv.ParseUint(data, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(value), nil
}

func StringToInt16(data string) (int16, error) {
	value, err := strconv.ParseInt(data, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(value), nil
}

func StringToUint16(data string) (uint16, error) {
	value, err := strconv.ParseUint(data, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(value), nil
}

func StringToInt32(data string) (int32, error) {
	value, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func StringToUint32(data string) (uint32, error) {
	value, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(value), nil
}

func StringToInt64(data string) (int64, error) {
	value, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func StringToUint64(data string) (uint64, error) {
	value, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func StringToFloat32(data string) (float32, error) {
	value, err := strconv.ParseFloat(data, 32)
	if err != nil {
		return 0, err
	}
	return float32(value), nil
}

func StringToFloat64(data string) (float64, error) {
	value, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func DecodeBytesToString(typeName string, data []byte, order binary.ByteOrder, encoding string) (string, error) {
	var value interface{}
	var err error
	switch typeName {
	case "int8":
		value, err = BytesToInt8(data, order)
	case "uint8":
		value, err = BytesToUint8(data, order)
	case "int16":
		value, err = BytesToInt16(data, order)
	case "uint16":
		value, err = BytesToUint16(data, order)
	case "int32":
		value, err = BytesToInt32(data, order)
	case "uint32":
		value, err = BytesToUint32(data, order)
	case "int64":
		value, err = BytesToInt64(data, order)
	case "uint64":
		value, err = BytesToUint64(data, order)
	case "float32":
		value, err = BytesToFloat32(data, order)
	case "float64":
		value, err = BytesToFloat64(data, order)
	default:
		return BytesToUtf8String(data, encoding)
	}

	if err != nil {
		return string(data), err
	}

	return fmt.Sprint(value), nil
}

func EncodeStringToBytes(typeName string, data string, order binary.ByteOrder, encoding string) ([]byte, error) {
	var value interface{}
	var err error
	switch typeName {
	case "int8":
		value, err = StringToInt8(data)
	case "uint8":
		value, err = StringToUint8(data)
	case "int16":
		value, err = StringToInt16(data)
	case "uint16":
		value, err = StringToUint16(data)
	case "int32":
		value, err = StringToInt32(data)
	case "uint32":
		value, err = StringToUint32(data)
	case "int64":
		value, err = StringToInt64(data)
	case "uint64":
		value, err = StringToUint64(data)
	case "float32":
		value, err = StringToFloat32(data)
	case "float64":
		value, err = StringToFloat64(data)
	default:
		return Utf8StringToBytes(data, encoding)
	}

	if err != nil {
		return []byte(data), err
	}

	return EncodeValueToBytes(value, order)
}
