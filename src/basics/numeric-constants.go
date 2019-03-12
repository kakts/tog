package main

import "fmt"

// 数値の定数は　高精度な"値"
const (
	// 1bitを100回左シフトすることで巨大数を作る
	Big = 1 << 100

	// 右シフトして2となる
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}