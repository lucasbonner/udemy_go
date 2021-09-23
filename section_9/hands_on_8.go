//create program that uses switch wt no switch expression specified
package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false big boy")
	case true:
		fmt.Println("yea that true")
	default:
		fmt.Println("defaulT")
	}
}
