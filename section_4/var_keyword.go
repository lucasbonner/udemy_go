package main

import "fmt"

var y = 43 //okay outside function body, scope is 'package level'
var z int  //declared variable z, and that it is of type int, assigns the zero value of type int

func main() {
	x := 42 //okay in function body
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
