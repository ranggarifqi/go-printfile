package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
