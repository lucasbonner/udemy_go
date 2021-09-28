///using composite literal:
//create an array which holds 5 values of type int
///assign values to each index
//range over array and print values out
//use format printing
//print out the type of the array

package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", arr)
}
