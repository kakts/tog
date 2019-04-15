package main

import "fmt"

func main() {
	// mapの生成
	m := make(map[string]int)

	// 要素の挿入や更新
	m["Answer"] = 42
	fmt.Println("The Value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value: ", m["Answer"])

	// 要素の削除
	delete(m, "Answer")
	fmt.Println("The Value: ", m["Answer"])

	// キーに対する要素が存在するかは2つ目の値で確認
	v, ok := m["Answer"]
	fmt.Println("The Value: ", v, "Present?", ok)
}