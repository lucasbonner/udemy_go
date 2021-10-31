package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //& gives address in memory

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) //*int means pointer to an int

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) //* operator, dereferences an address
	//b points to address of a, *b gives value stored at that address

	c := a
	fmt.Println(c)
	fmt.Println("Address for a is:", &a, "Address for c is", &c) //these addresses are different
}
