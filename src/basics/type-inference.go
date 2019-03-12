package main

import "fmt"

func main() {
	// 明示的な型変換なしで変数を宣言する場合　:= or var = 
	// 変数の型は右側の変数から型推論される

	v := 42.1 // change me
	fmt.Printf("v is of type %T\n", v)
}