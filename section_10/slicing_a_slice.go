package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1:])  //[5, 7, 8, 42]
	fmt.Println(x[1:3]) //[5, 7]

	for i := 0; i <= len(x)-1; i++ {
		fmt.Println(x[i])
	}
}
