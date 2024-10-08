package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Bob() *User {
	// 関数ないの変数userのポインタを関数外へ返す
	user := User{
		Name: "Bob",
		Age:  18,
	}
	return &user
}

func main() {
	user := Bob()
	user.Name = "Alice"
	user.Age = 20
	fmt.Println(user) // &{Alice 20}
}
