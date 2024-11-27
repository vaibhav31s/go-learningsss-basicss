package main

import "fmt"

func main() {
	fmt.Println("Interfaces in Go")

	s := square{length: 4}
	c := circle{reaius: 5}

	printShapeInof(s)
	printShapeInof(c)
}

type shape interface {
	area() float64
	circumference() float64
}

type square struct {
	length float64
}

type circle struct {
	reaius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return 4 * s.length
}

func (c circle) area() float64 {
	return 3.14 * c.reaius * c.reaius
}

func (c circle) circumference() float64 {
	return 2 * 3.14 * c.reaius
}

func printShapeInof(s shape) {
	fmt.Println("Area is: ", s.area())
	fmt.Println("Circumference is: ", s.circumference())
}
