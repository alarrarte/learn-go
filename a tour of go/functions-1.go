package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func take(x, y int) int {
	return x - y
}

func main() {

	fmt.Println(add(42, 70))
	fmt.Println(take(42, 70))
}
