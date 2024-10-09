package main

import "reflect"

type Foo struct {
	Name string
}

func main() {
	var foo Foo
	println(reflect.TypeOf(foo))
	foo.Name = "foo"
}
