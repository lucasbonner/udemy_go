//Create TYPED and UNTYPED constants
package main

import "fmt"

//untyped constants
const (
	a = "dog"
	b = "cat"
	c = 42
)

//typed constants
const (
	d int     = 28
	e string  = "cow"
	f float64 = 2.22
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
