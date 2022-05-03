package transform

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func BytesToUtf8String(data []byte, encoding string) (string, error) {
	if encoding == "gbk" {
		ud, err := GbkToUtf8(data)
		if err != nil {
			return string(data), err
		}
		return string(ud), nil
	}
	return string(data), nil
}

func Utf8StringToBytes(str string, encoding string) ([]byte, error) {
	if encoding == "gbk" {
		data, err := Utf8ToGbk([]byte(str))
		if err != nil {
			t := []byte(str)
			l := len(t)
			for i := l - 1; i > 0; i-- {
				if data, err := Utf8ToGbk(t[:i]); err == nil {
					return data, nil
				}
			}
			return []byte(str), err
		}
		return data, nil
	}
	return []byte(str), nil
}
