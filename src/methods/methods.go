package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 型にメソッドを定義できる
// funcとメソッド名の間に自身の引数リストで表現する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}