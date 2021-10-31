/*
create a func which returns a func
assign the returned func to a variable
call the returned func
*/
package main

import "fmt"

func main() {

	x := hello()
	fmt.Println(x())
}

func hello() func() int {
	return func() int {
		return 25
	}
}
