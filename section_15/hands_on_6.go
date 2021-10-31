/*
Build and use an anonymous func
*/
package main

import "fmt"

func main() {
	func(s string) {
		fmt.Println("I am anonymous, and I was given:", s)
	}("dog")
}
