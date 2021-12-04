/*
create a type person struct
attach a method speak to type pointer to a person
	*person
create a type human interface
	to implicitly implement the interface a human must have speak method
create func "saySomething"
	have it take in a human as a parameter
	have it call the speak method
show the following in code
	you can pass type *person into saySomething
	you CANNOT pass type person into saySomething
*/
package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return p.first + " is speaking"
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	lucas := person{
		first: "lucas",
		last:  "boner",
	}

	saySomething(&lucas)
}
