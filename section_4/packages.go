package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, playground", 42, true)
	fmt.Println(n)   //number of bytes written
	fmt.Println(err) //error
}
