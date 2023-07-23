package multireader

import (
	"fmt"
	"io"
	"log"
	"os"
)

type MultiReader struct {
	reader io.Reader
	file1  os.File
	file2  os.File
}

func NewMultiReader(r io.Reader, f1 *os.File, f2 *os.File) *MultiReader {
	return &MultiReader{
		reader: r,
		file1:  *f1,
		file2:  *f2,
	}
}

func (mr *MultiReader) ReadTwoFilesWithMultiReader() {
	info, _ := mr.file1.Stat()
	bufferSizeFile1 := info.Size()
	info, _ = mr.file2.Stat()
	bufferSizeFile2 := info.Size()

	// слайс байт для буффера, в который ридер запишет текст из файлов
	buffer := make([]byte, bufferSizeFile1+bufferSizeFile2)

	mr.reader.Read(buffer[:bufferSizeFile1])
	mr.reader.Read(buffer[bufferSizeFile1:])
	fmt.Println(string(buffer))
}

func (mr *MultiReader) ReadTwoFilesWithBiggerBuffer() {
	biggerBuffer := make([]byte, 10)
	for {
		readBytes, err := io.ReadFull(mr.reader, biggerBuffer)
		if err != nil {
			fmt.Printf("Bytes read: %d, buffer cap: %d\n", readBytes, cap(biggerBuffer))
			log.Fatal(err)
		}
		fmt.Println(readBytes)
	}

	//fmt.Println(string(biggerBuffer))
}
