package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 7 = ", mySum(4, 7))
}

func mycapital(xi string) string {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
