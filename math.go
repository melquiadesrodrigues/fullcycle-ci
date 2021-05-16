package main

import "fmt"

func main() {
	fmt.Println(Sum(10, 10))
}

func sum(a int, b int) int {
	return a + b
}