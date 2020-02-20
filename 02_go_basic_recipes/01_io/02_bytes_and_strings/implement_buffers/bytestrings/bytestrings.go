package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

func Buffer(rawString string) *bytes.Buffer {

	// encode string into raw bytes
	rawBytes := []byte(rawString)

	// Can create buffer from raw bytes
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// Alternatively
	b = bytes.NewBuffer(rawBytes)

	// or... avoid initial byte array
	b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
