package main

import "fmt"

var x bool //zero value is false

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
