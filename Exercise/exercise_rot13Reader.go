package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read (b []byte) (int, error) {
	n,err := rot.r.Read(b)
	for i,v := range b {
		if err != nil {
			return 0,err
		}
		if 	('A' <= v && v <= 'M') || ('a' <= v && v <= 'm'){
			b[i] += 13
		}
		if ('N' <= v && v <= 'Z') || ('n' <= v && v <= 'z'){
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}