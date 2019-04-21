package main

import (
	"fmt"
	"math"
)

// ポインタレシーバを使う2つの理由
// 1: メソッドがレシーバがさす先の変数を変更するため
// 2: メソッドの呼び出しごとに変数のコピーを避けるため　特にレシーバが大きな構造体の場合に効率的


type Vertex struct {
	X, Y float64
}
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}