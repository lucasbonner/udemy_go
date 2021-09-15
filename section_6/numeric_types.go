package main

import "fmt"

var x int
var y float64

func main() {
	x = 42       //int
	y = 42.34544 //float64

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
