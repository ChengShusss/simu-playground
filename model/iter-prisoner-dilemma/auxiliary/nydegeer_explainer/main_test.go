package main

import (
	"fmt"
	"testing"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func TestString(t *testing.T) {
	points := []Point{{1, 2}, {3, 4}}
	fmt.Println(points) // 输出：[(1, 2) (3, 4)]

	a := ActCooperate
	fmt.Println(a)
	fmt.Println(&a)
}

func TestNewCase(t *testing.T) {
	c := NewCase(63)
	fmt.Printf("%v\n", c)

	for _, line := range c.Format2Lines() {
		fmt.Println(line)
	}
}
