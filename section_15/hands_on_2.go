/*
create a func with the identifier foo that
	takes in a variadic parameter of type int
	pass in a value of type {}int into your func (unfurl the []int)
	returns the sum of all values of type int passed in

create a func with the identifier bar that
	takes in a parameter of type []int
	returns the sum of all values of type int passed in
*/
package main

import "fmt"

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	return foo(x...)
}

func main() {
	xi := []int{1, 2, 3, 4, 5}
	xii := []int{5, 6, 7, 8607, 7}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xii))
}
