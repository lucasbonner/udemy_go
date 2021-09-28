/*
start with this slice:
x:=[]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

append 52 to slice

print out slice

in ONE STATEMENT append to slice:
53, 54, 55

print out the slice

append the slice to this slice
y := []int{56, 57, 58, 59, 60}

print out slice
*/

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)

	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
