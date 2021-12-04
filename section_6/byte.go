package main

import "fmt"

func main() {
	s := "Hello, moto"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) //slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i, v := range bs {
		fmt.Println(i, v)
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
