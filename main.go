package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func showName(user *User) {
	fmt.Println(user.Name)
}

func main() {
	user := User{
		Name: "Jane",
		Age:  20,
	}
	showName(&user) // Jane
}
