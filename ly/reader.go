package ly

import (
	"io"
	"os"
	"strings"
)

type MyReader struct{}

func (m MyReader) Read(b []byte) (n int, err error) {
	n, err = 1, nil
	b[0] = 'A'
	return
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = b + 13
	case 'M' < b && b <= 'Z':
		b = b - 13
	case 'a' <= b && b <= 'm':
		b = b + 13
	case 'm' < b && b <= 'z':
		b = b - 13
	}
	return b
}

func (mr rot13Reader) Read(b []byte) (n int, err error) {
	n, e := mr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, e
}

func R13Read() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
