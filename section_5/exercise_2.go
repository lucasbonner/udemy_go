//var to declare three variables
//package level scope
//do not assign values to variables
//use following identifiers

//identifier "x" type int
//identifier "y" type string
//identifier "z" type bool

//in func main:
//print out values for each identifier
//compiler assigned values to the variables
//what are these values called -> zero values

package main

import "fmt"

var x int    // 0
var y string // ''
var z bool   // false

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
