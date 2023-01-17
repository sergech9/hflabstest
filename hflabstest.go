package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	file, _ := os.Open("QAtest.txt")
	count, _ := lineCounter(file)
	fmt.Println(count)
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'a'} // проблема с кириллической «а» таится тут

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
