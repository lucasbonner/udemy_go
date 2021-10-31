package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//interface:
type human interface {
	//any other type with speak() method is also of type human
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrr", h.(secretAgent).first)
	}

}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(sa2)
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
