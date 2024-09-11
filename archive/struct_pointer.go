package main

import "fmt"

type User2 struct {
	id   string
	name string
	age  int
}

/*
 * ポインタ型で使い回すことで、メモリに格納された値を使い回し続けるので
 * メモリ消費量を節約可能
 * 構造体など複数の値の場合は重くなるので、ポインタ型にするのが一般的
 */
func main() {
	// &をつけてポインタ型に
	user := &User2{
		id:   "1",
		name: "John Lennon",
		age:  30,
	}

	fmt.Printf("%s\n", user.name)
	user.name = "Paul"
	fmt.Printf("%s\n", user.name)

}
