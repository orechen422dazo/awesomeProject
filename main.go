package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	fmt.Println(len(slice)) // 出力: 5
	fmt.Println(cap(slice)) // 出力: 10
}
