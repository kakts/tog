package main

import "fmt"

// 変数Tのポインタは *T型で　ゼロ値はnil
func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // ポインタを通じてiを読み取る
	*p = 21 // ポインタを通じてiをセットする
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}