package main

import (
	"fmt"
	"math"
)

// START IMPL OMIT

type rect struct {
	width, height float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	Measure(rect{2, 4})
	Measure(circle{3})
}

// END IMPL OMIT

// START INTERFACE OMIT

type Geometry interface {
	Area() float64
}

func Measure(g Geometry) {
	fmt.Println("Area:", g.Area())
}

// END INTERFACE OMIT
