package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fn := os.Args[1]
	fmt.Println("Trying to open file with name:", fn)

	f, err := os.Open(fn)

	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}

	lw := logWriter{}

	io.Copy(lw, f)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
