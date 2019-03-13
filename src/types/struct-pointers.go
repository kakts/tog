package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// structのフィールドは、structのポインタを通じてアクセスすることもできる
	p := &v

	// structのポインタpのフィールドにアクセスする場合　(*p).Xのようにかけるが
	// goではp.Xとかける
	p.X = 1e9
	fmt.Println(v)
}