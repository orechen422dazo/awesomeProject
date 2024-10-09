package main

import "fmt"

var v interface{}

func PrintDetail(v interface{}) {
	switch t := v.(type) {
	case int, int32, int64:
		fmt.Print("int/int32/int64 data type:", t)
	case string:
		fmt.Print("string", t)
	default:
		fmt.Print("unknown type")
	}
}

func main() {
	v = 100
	PrintDetail(v)
	fmt.Println()
}
