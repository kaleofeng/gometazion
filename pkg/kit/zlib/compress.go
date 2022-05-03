package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Compress(data []byte) ([]byte, error) {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	w.Close()
	return in.Bytes(), nil
}

func Uncompress(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var out bytes.Buffer
	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
