/*
Use the "defer" keyword to show that a deferred func runs after the func
containing it exits
*/

package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("I'm first, but then again I'm deferred")
}

func bar() {
	fmt.Println("I didn't want to go first!!!")
}
