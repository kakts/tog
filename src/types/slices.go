package main

import "fmt"

func main() {

	// 配列
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// []T は型Tのスライスを表す
	var s []int = primes[1:4]
	fmt.Println(s)
}