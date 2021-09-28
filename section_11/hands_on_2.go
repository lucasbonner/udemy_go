package main

import "fmt"

func main() {
	slice := []int{3, 28, 6, 7, 9}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)
}
