package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, " + strconv.Itoa(add(2, 3)))
}
func add(a int, b int) int {
	return a + b
}
