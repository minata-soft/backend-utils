package backend_utils

import (
	"bytes"
	"io"
)

var IOConvert = ioStruct{}

type ioStruct struct {
}

func (ioStruct) ReadCloserToString(b io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(b)
	respBytes := buf.String()

	respString := string(respBytes)

	return respString
}
