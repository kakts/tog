package main

import "fmt"

/**
 * チャネルはバッファとして使える
 * 
 * バッファを持つチャネルを初期化するには　makeの2つ目の引数にバッファ長を与える
 * 
 * バッファが詰まったときは、チャネルへの送信をブロックする
 * バッファが空の時には、チャネルの受信をブロックする
 */
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4


	fmt.Println(<-ch)
	fmt.Println(<-ch)
}