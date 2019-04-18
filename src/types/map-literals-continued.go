package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapに渡すトップレベルの型が単純な型名の場合はリテラルの要素から推定できるので、
// 型名を省略できる Vertex{...}と書かなくて良い
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408}
}

func main() {
	fmt.Println(m)
}