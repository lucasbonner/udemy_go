/*
assign a func to a variable, then call that func
*/
package main

import "fmt"

func main() {
	v := func(x int) {
		for i := x; i > 0; i-- {
			fmt.Println("I will print", x, "times")
		}
	}

	v(4)
	v(10)
}
