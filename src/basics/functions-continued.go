package main

import "fmt"

// 2つ以上の引数が同じ型の場合は　最後の型だけ記述すればよい
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}