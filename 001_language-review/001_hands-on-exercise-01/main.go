package main

import (
	"fmt"
	"math"
)

type square struct {
	s float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.s * s.s
}

func (c circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

type shape interface {
	area() float64
}

func main() {
	c := circle{r: 2}
	s := square{s: 4}

	info(c)
	info(s)
}

func info(s shape) {
	fmt.Println(s.area())
}
