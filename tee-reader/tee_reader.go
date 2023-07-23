package tee_reader

import (
	"fmt"
	"io"
	"strings"
)

type TeeReader struct {
	reader io.Reader
}

func NewTeeReader(r io.Reader) *TeeReader {
	return &TeeReader{
		reader: r,
	}
}

func (cr *TeeReader) Read() {
	buffer := make([]byte, 5)
	builder := strings.Builder{}
	for {
		n, err := cr.reader.Read(buffer)
		builder.Write(buffer[:n])
		fmt.Printf("Bytes read: %d\n", n)
		if err == io.EOF {
			break
		}
	}
	fmt.Println(builder.String())
}
