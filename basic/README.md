# Basic

### var, const, iota
data type check method `reflect.TypeOf(num)`
```go
package main

import (
	"log"
	"reflect"
)

func main() {
	num := 1
	log.Print(num)
	log.Println("type:", reflect.TypeOf(num))
	
	t := "ffff"
	log.Println(t)
	log.Println("type:", reflect.TypeOf(t))
	
	b := true
	log.Println(b)
	log.Println("type:", reflect.TypeOf(b))
	
	arr := [3]int{1, 2, 3}
	log.Println(arr)
	log.Println("type:", reflect.TypeOf(arr))
	
	sl := []int{1, 2, 3}
	log.Println(sl)
	log.Println("type:", reflect.TypeOf(sl))
	
	m := map[string]int{"a": 1, "b": 2}
	log.Println(m)
	log.Println("type:", reflect.TypeOf(m))
}
```

### if
x == 20 OK
```go
package main

func main() {
	x := 20
	if x == 30 {
		println("x is 10")
	}
	if x == 20 {
		println("x is 20")
	}
}
```

### for slice map
loop example
```go
package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	println("-------------- slice --------------")

	// slice
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		println(i, v)
	}

	println("------------- map --------------")
	// map
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v)
	}
}
```

### type
data type naming
```go
package main

type Foo string

func main() {
	var foo Foo
	foo = "bar"
	println(foo)
}
```

case string
```go
package main

import "reflect"

type Foo string

func main() {
	var foo Foo
	foo = "bar"
	println(foo)
	println(reflect.TypeOf(foo))
	s := string(foo)
	println(s)
	println(reflect.TypeOf(s))
}
```

### struct
data type Foo
```go
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
```