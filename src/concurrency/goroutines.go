package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
/**
 * goroutineは、Goのランタイムに管理される軽量なスレッド
 * 
 * go f(x, y, z)とかけば新しいgoroutineが実行される
 * f x y zの評価は、実行元のgoroutineで実行される
 * fの実行は新しいgoroutineで実行される
 * 
 * goroutineでは同じアドレス空間で実行されるため
 * 共有メモリへのアクセスは必ず同期する必要がある
 * 
 * syncパッケージは同期する際に役に立つ方法を提供するが、別の方法があるためそれほど必要ではない
 */
func main() {
	go say("world")
	say("hello")
}