/*
Closure is when we have “enclosed” the scope of a variable in some
code block. For this hands-on exercise, create a func which “encloses” the
scope of a variable:
*/
package main

import "fmt"

func main() {
	fmt.Println(closure()())
}

func closure() func() int {
	x := 3
	return func() int {
		x++
		return x
	}
}
