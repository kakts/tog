package main

import "fmt"

func main() {
	var i interface{} = "hello"

	/**
	 * 型アサーション
	 * インターフェースの値の基になる具体的な値を利用する手段を提供
	 * 
	 * t := i.(T)
	 * インタフェースの値iが具体的な型Tを保持し、
	 * 基になるTの値を変数tに代入することを主張する
	 * 
	 * iがTを保持していない場合この文はpanicを引き起こす
	 */
	s := i.(string)
	fmt.Println(s)

	/**
	 * インターフェースの値が特定の型を保持しているかどうかをテストするため
	 * 型アサーションは2つの値（基になる値とアサーション成功したかどうかのbool)
	 * を返す
	 */
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f, ok = i.(float64) // panic
	fmt.Println(f, ok)
}