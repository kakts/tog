package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// 初期値なしの場合、ゼロ値が与えられる
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}