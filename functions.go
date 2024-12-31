package main

import "fmt"

func add(x int, y int) { // helper function
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}