package main

import "fmt"

// constキーワードにより定数を宣言する
// 定数は :=をつかって宣言できない
const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}