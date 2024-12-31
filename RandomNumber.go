package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myBool bool = true
	if myBool {
		fmt.Println("My favorite number is", rand.Intn(10))
	}
}
