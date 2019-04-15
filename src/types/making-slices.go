package main

import "fmt"


func main() {
	// makeはゼロ化された配列を割り当て、その配列をさすスライスを返す
	a := make([]int, 5)
	printSlice("a", a)

	// 3 つめの引数に容量を指定できる
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}