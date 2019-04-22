package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

/**
 * T はインターフェースIを実装する
 * implementsキーワードは必要ない
 */
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}