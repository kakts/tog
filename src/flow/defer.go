package main

import "fmt"

func main() {
	// 呼び出し元の関数の終わり(returnする)まで遅延させる
	// deferに渡した関数の引数は　すぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで
	// 実行されない
	defer fmt.Println("world")

	fmt.Println("Hello")
	return
}