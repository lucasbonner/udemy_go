package main

import "fmt"

func main() {
	//right when func main exits, defer funcs run
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
