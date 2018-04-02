package httpReq

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GzipEncode(orgData []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)

	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(orgData)

	if err != nil {
		writer.Close()
		return out, err
	}

	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func GzipDecode(orgData []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(orgData))

	var out []byte

	if err != nil {

		return out, err
	}

	defer reader.Close()

	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return out, err
	}

	return buffer, err
}
