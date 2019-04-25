package main

import (
	"fmt"
)

/**
 * フィボナッチ数列計算
 * 
 * チャネル送信側はこれ以上送信することがないことを示すためcloseできる
 */
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	fmt.Println("Finish fibonacci")
	close(c)
}

/**
 * チャネル送信側はこれ以上送信することがないことを示すためcloseできる
 * 
 * 受けては受信の式に2つめのパラメータを割り当てることで
 * そのチャネルがクローズされているかを確認できる
 * 
 * v, ok := <- ch
 * 受信する値がない、もしくはチャネルが閉じているならok = falseになる
 * 
 * 注意　送り手のチャネルだけをcloseする　受けてはcloseしてはいけない　するとpanicになる
 * 
 * チャネルは、ファイルとは異なり、通常は、closeする必要はありません。 
 * closeするのは、これ以上値が来ないことを受け手が知る必要があるときにだけです。 例えば、 range ループを終了するという場合です。
 */
func main() {
	// バッファ付きチャンネルを作成する
	c := make(chan int, 10)

	go fibonacci(cap(c), c)
	// fmt.Println("Finish fibonacci")
	
	// チャネルが閉じられるまでチャネルから値を繰り返し受信し続ける
	for i := range c {
		fmt.Println(i)
	}
}