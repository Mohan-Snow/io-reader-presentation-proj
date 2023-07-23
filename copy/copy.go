package copy

import (
	"fmt"
	"io"
	"os"
)

type Copier struct {
	reader io.Reader
}

func NewCopier(r io.Reader) *Copier {
	return &Copier{
		reader: r,
	}
}

func (c *Copier) CopyBytesFromBufferToFile() {
	dst, _ := os.Create("./misc/temp")
	defer dst.Close()
	readBytes, err := io.Copy(dst, c.reader)
	if err != nil {
		fmt.Println("Error while copying to the file")
	}
	fmt.Printf("The number of copied bytes to file: %d\n", readBytes)

	buffer := make([]byte, 5)
	read, _ := c.reader.Read(buffer)
	fmt.Printf("Bytes in reader: %d\n", read)
}
