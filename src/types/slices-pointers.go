package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// スライスはどんなデータも格納しておらず、単に元の配列の部分列を指し示す
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスの要素を変更すると、同じ元となる配列の対応する要素も変更される
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}