package main

import "fmt"

type I interface {
	M()
}

/**
 * 呼び出す具体的なメソッドを示す型がインターフェースのタプル内に存在しない
 * nilインタフェースのメソッドを呼び出すとランタイムエラーになる
 */
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}