package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func createRotMap() map[string]string {
	var rotMap map[string]string
	src := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rot := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	ss := strings.Split(src, "")
	rr := strings.Split(rot, "")
	for i := 0; i < len(ss); i++ {
		rotMap[ss[i]] = rr[i]
	}
	fmt.Println(rotMap)
	return rotMap

}

// Read is
func (r rot13Reader) Read(p []byte) (n int, err error) {
	createRotMap()
	/*
		rotMap := createRotMap()
		fmt.Println(rotMap)
			n, err = r.r.Read(p)
			for i := 0; i < n; i++ {
				v, ok := rotMap[string(p[i])]
				fmt.Println(ok)
				fmt.Println(v)
			}
	*/
	return len(p), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
