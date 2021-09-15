//create own type, uinderlying type int
//create variable of new type with identifier x using var

//in func main
//print out value of variable 'x'
//print out the type of variable 'x'
//assign 42 to variable 'x' using '=' operator
//print out value of variable x

package main

import "fmt"

type gremlin int

var x gremlin

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
