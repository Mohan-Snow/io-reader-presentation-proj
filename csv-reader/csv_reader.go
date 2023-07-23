package csv_reader

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvReader struct {
	reader csv.Reader
}

func NewCsvReader(r csv.Reader) *CsvReader {
	return &CsvReader{
		reader: r,
	}
}

func (cr *CsvReader) ReadCsvFile() {
	for {
		line, err := cr.reader.Read()
		fmt.Printf("%s\n", line)
		if err == io.EOF {
			break
		}
	}
}
