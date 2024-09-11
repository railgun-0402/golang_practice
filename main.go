package main

import (
	"flag"
	"fmt"
	"os"
)

type User struct {
	id   string
	name string
	age  int
}

func (u User) Greet() {
	fmt.Printf("I`m %s\n", u.name)

	fmt.Printf("%s\n", u.name)
	// => John Lennon
	u.name = "Paul McCartney"
	fmt.Printf("%s\n", u.name)
	// => Paul McCartney
}

func main() {
	user := User{
		id:   "1",
		name: "John Lennon",
		age:  30,
	}
	user.Greet()

	// 入力値を返す
	flag.Parse()
	arg := flag.Arg(0)
	// 関数の返り値として返す
	msg := fmt.Sprintf("Hello %s\n", arg)

	f, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	// defer f.Close()

	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}

	fmt.Println("Done!")
}
