package main

import "fmt"

// START OMIT
type Car struct {
	x, y int
}

func (c *Car) moveForward(steps int) {
	c.y += steps
}

func (c Car) moveBackwards(steps int) {
	c.y -= steps
}

func main() {
	var c Car
	c.moveForward(3)
	c.moveBackwards(2)
	fmt.Println(c.x, c.y)
}

// END OMIT
