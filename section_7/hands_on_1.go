package main

//Write a program that prints a number in decimal, binary, and hex

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)
}
