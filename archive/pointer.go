package main

import "fmt"

// stringのポインタ型で引数を受け取る
func sayHello(name *string) {
	fmt.Printf("Hello, %s\n", *name)
}

func main() {
	a := "World" // string型は*string型に仮引数として、渡せない
	b := &a      // ポイント型同士は引き渡し可能
	sayHello(b)
}
