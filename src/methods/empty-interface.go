package main

import "fmt"

func main() {
	/**
	 * 空のインターフェースは任意の型の値を保持できる
	 */
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	// fmt.Printfはinterface{}型の任意の数の引数を受け取る
	fmt.Printf("(%v, %T)\n", i, i)
}