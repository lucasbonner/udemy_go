package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x) //[0 0 0 0 0]
	x[3] = 42
	fmt.Println(x)      //[0 0 0 42 0]
	fmt.Println(len(x)) //5
}
