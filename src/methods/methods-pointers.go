package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// ポインタレシーバでメソッドを宣言できる
// このメソッドは、レシーバがさす変数を変更できる
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 変数レシーババージョン　元の値は変更されない
func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	v.Scale2(10)
	fmt.Println(v.Abs())
	v.Scale(10)

	fmt.Println(v.Abs())
}