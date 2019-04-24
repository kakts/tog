package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/** wrong
 * 
func (rot rot13Reader) Read(b []byte) (n int, err error) {
	for {
		rotN, rotErr := rot.r.Read(b)

		fmt.Printf("----")
		fmt.Printf("b[:n] = %q\n", b[:n])
		for n, err = 0, nil; n < rotN; n++ {
			b[n] = b[n] + 13
		}
		
		if rotErr == io.EOF {
			break
		}
	}
	return
}
*/

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	// bに対して読み込む
	n, err = rot.r.Read(b)
	if err == nil {
		for i, v := range b {
			switch {
			case v >= 'A' && v <= 'Z':
				b[i] = (v - 'A' + 13) % 26 + 'A'
			case v >= 'a' && v <= 'z':
				b[i] = (v - 'a' + 13) % 26 + 'a'
			}
		}
	 }
	 return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}