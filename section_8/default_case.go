package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 4):
		fmt.Println("prints")
		fallthrough
	case (7 == 9):
		fmt.Println("also true, does it print?")
		fallthrough
	case (42 == 3):
		fmt.Println("this is not a true statement")
	default:
		fmt.Println("this is default")
	}
}
