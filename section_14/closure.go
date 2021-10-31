package main

import "fmt"

//package level scope: x (global)
// var x int

func main() {
	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y) won't work, y is limited to code block

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		//don't have to pass x in, it is within scope of function
		x++
		return x
	}
}
