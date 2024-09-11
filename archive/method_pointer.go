package main

import "fmt"

type User struct {
	id   string
	name string
	age  int
}

func (u *User) IncrementAge() {
	u.age++
}

// IncrementAgeNoPointer 引数がポインタ型かどうか分けて検証
func (u User) IncrementAgeNoPointer() {
	u.age++
}

func main() {
	// &をつけてポインタ型に
	user := &User{
		id:   "1",
		name: "John Lennon",
		age:  30,
	}

	// 引数Userがポインタではないので、メソッド実行の際に値がコピーされて渡される
	// コピーした場合は、メソッド内の更新はメソッド外では反映されない
	user.IncrementAgeNoPointer()
	fmt.Println(user.age) // 30

	// ポインタ型の場合は、レシーバを定義するので更新結果がメソッド外でも反映される
	user.IncrementAge()
	fmt.Println(user.age) // 31

}
