/*
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle

*/
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius int
}

type square struct {
	base   int
	height int
}

func (c circle) area() float64 {
	a := math.Pow(float64(c.radius), 2) * math.Pi
	return a
}

func (s square) area() float64 {
	a := s.base * s.height
	return float64(a)
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println(s.(square).area())
	case circle:
		fmt.Println(s.(circle).area())
	}
}

func main() {
	circ := circle{
		radius: 10,
	}

	squar := square{
		base:   15,
		height: 15,
	}

	info(circ)
	info(squar)
}
