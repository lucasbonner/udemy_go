/// func (r receiver) identifier(parameters) (return(s)) {...}
package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

//everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hell from woo, ", st)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b
}
