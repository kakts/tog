package main

import (
	"fmt"
	"math"
)

// 任意の型にもメソッドを定義できる
type MyFloat float64

// メソッドの宣言は　レシーバ型が同じパッケージにある必要がある
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}