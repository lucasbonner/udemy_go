package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}
	f()
	m := func(x int) {
		fmt.Println("The meaning of life:", x)
	}
	m(42)
}
