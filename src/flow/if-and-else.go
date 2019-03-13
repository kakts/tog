package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if ステートメントで宣言された変数はelseブロック内でも使える
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%q >= %q\n", v, lim)
	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}