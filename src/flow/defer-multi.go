package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// deferは複数指定可能　呼び出しはスタックされる
		// last in first out
		defer fmt.Println(i)
	}

	fmt.Println("done")
}