//use previous code

//package level scope, assign following values to three variables

//for x assign 42
//for y assign "James Bond"
//for z assign true

//in func main
//use fmt.Sprintf to print all of the VALUES to one string
//assign return value of type string using short decleration operator
//to "s"

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
