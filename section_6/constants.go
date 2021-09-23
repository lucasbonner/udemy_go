package main

import "fmt"

const (
	a int     = 42           //typed constant
	b float64 = 42.78        //untyped constant
	c string  = "James Bond" //untyped constant
)

// const a = 42 untyped constant
// const b = 42.78 untyped constant
// const c = "James Bond" untyped constant

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
