package main

import "fmt"

/**
 * スライスは　長さlength と 容量capacityを持つ
 * 長さ　含まれる要素の数
 * 容量　スライスの最初の要素から数えて、元となる配列の要素数
 * 長さ　len(S)  容量 cap(S) で取得可能
 */
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 0の長さにスライスする
	s = s[:0]
	printSlice(s)

	// 長さを伸ばす
	s = s[:4]
	printSlice(s)

	// 最初の2つの要素を削除する
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap=%d %v\n", len(s), cap(s), s)
}