package main

import "fmt"

func add(x int, y int) string { // helper function
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}