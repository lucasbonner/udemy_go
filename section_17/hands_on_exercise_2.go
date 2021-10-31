/*
	create a person struct
	create a func called "changeMe" with a *person as a parameter
		change the value stored at the *person address
			important
				to dereference a struct, use (*value) field
				as an exception if the type of x is a named pointer type and (*x)
				is a valid selector expression denoting a field(but not a method
				x.f is shorthand for (*x).f.

	create a value of type person
		print out the value
	in func main
		call changeMe
	in func main
		print out the value
*/
package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) *person {
	(*p).first = "yuck"
	return p
}

func main() {
	p1 := person{
		first: "lucas",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
