/*
a callback is when we pass a func into a func as an argument
	pass a func into a func as an argument
*/
package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(sum, 5))
}

func sumOfMultiples(f func([]int) int, x int) int {
	result := []int{}
	for i := 1; i < x; i++ {
		result = append(result, i)
	}
	return f(result)
}

func sum(sli []int) int {
	sum := 0
	for i := 0; i < len(sli); i++ {
		sum += sli[i]
	}

	return sum
}
