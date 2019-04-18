package main

import "fmt"

func fibonacci() func() int {
	count := 0
	x1 := 0
	x2 := 1
	y := 0
	return func() int {
		// fmt.Println(count)
		if count < 2 {
			res := count
			count += 1
			return res
		}
		y = x1 + x2
		x1 = x2
		x2 = y
		count += 1
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}