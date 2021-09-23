//Create a program that shows the "if statement" in action, else if, and else

package main

import "fmt"

func main() {
	str := "dog"

	if str == "cat" {
		fmt.Println("meow")
	} else if str == "dog" {
		fmt.Println("woof")
	} else {
		fmt.Println("yee")
	}
}
