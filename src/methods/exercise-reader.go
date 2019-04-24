package main

import (
	"golang.org/x/tour/reader"
	"fmt"
)
type MyReader struct {}

func (mr MyReader) Read(rb []byte) (n int , e error) {
	for n, e = 0, nil; n < len(rb); n++ {
		rb[n] = 'A'
	}
	return 
}

func main() {
	reader.Validate(MyReader{})

	r := &MyReader{}
	b := make([]byte, 10)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}
}