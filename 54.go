package main

import (
	"fmt"
	"os"
)

// Reader is
type Reader interface {
	Read(b []byte) (n int, err error)
}

// Writer is
type Writer interface {
	Write(b []byte) (n int, err error)
}

// ReadWriter is
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}
