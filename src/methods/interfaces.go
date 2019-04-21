package main

import (
	"fmt"
	"math"
)

/**
 * インタフェース
 * メソッドのシグニチャの集まりで定義する
 */
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// MyFloat implements Abser
	a = f

	// &Vertex implements Abser
	a = &v

	// aがVertexだが、Abserを実装していない
	a = v

	fmt.Println(A.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}