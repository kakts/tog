package main

import (
	"fmt"
	"io"
	"strings"
)

/**
 * io.ReaderインタフェースはReadメソッドを持つ
 * func (T) Read(b []byte) (n int, err error)
 * 
 * Readはデータを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーを返す
 * 
 */
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}