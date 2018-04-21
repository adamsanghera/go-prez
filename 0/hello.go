package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("hello" + "world")
	if err != nil {
		panic(err)
	}

	n = 3

	fmt.Printf("Num bytes: %d", n)
}
