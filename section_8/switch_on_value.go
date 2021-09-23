package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do no":
		fmt.Println("miss money")
	case "M":
		fmt.Println("M")
	default:
		fmt.Println("this is default")
	}
}
