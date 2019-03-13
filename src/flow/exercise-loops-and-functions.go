package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z * z - x) / (2 * z)
		if z * z - x <= 0.01 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}