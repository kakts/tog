package main

import "fmt"

func main() {
	sum := 1
	// whileがないのでforを使う
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}