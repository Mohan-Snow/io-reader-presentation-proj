package common_reader

import (
	"fmt"
	"io"
	"strings"
)

type CommonReader struct {
	reader io.Reader
}

func NewCommonReader(r io.Reader) *CommonReader {
	return &CommonReader{
		reader: r,
	}
}

func (cr *CommonReader) ReadString() {
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
