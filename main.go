package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	r "io-package-presentation-proj/common-reader"
	c "io-package-presentation-proj/copy"
	cr "io-package-presentation-proj/csv-reader"
	mr "io-package-presentation-proj/multireader"
)

func main() {
	//read()
	readMultipleFiles()
	//copyBytesFromBufferToFile()
}

func read() {
	reader := r.NewCommonReader(bytes.NewReader([]byte("This is a sample text :)")))
	reader.ReadString()
}

func readMultipleFiles() {
	file1, err := os.Open("./misc/test1")
	defer file1.Close()
	if err != nil {
		fmt.Println(err)
	}
	file2, err := os.Open("./misc/test2")
	defer file2.Close()
	if err != nil {
		fmt.Println(err)
	}
	mReader := mr.NewMultiReader(io.MultiReader(file1, file2), file1, file2)

	mReader.ReadTwoFilesWithBiggerBuffer()
	//mReader.ReadTwoFilesWithMultiReader()
}

func copyBytesFromBufferToFile() {
	copier := c.NewCopier(strings.NewReader("Source string to be copied to file."))
	copier.CopyBytesFromBufferToFile()
}

func readWithTeeReader() {
	file2, err := os.Open("./misc/test2")
	defer file2.Close()
	if err != nil {
		fmt.Println(err)
	}
	//bufio.NewWriter()
	//tee_reader.NewTeeReader(io.TeeReader(file2))
}

func readFromCsvFile() {
	cf, err := os.Open("./misc/my-csv-file.csv")
	defer cf.Close()
	if err != nil {
		fmt.Println(err)
	}
	csvReader := csv.NewReader(cf)
	cr.NewCsvReader(*csvReader)
}
