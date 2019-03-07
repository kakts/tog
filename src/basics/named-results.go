package main

import "fmt"

// 戻り値となる変数に名前を付けれる
// but 短い関数で使用すべき　可読性が悪くなる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// returnになにも書かずに戻せる
	return
}

func main() {
	fmt.Println(split(17))
}