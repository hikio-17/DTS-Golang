package main

import (
	"fmt"
	"math"
	"project/helpers"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}
	_ = r1

	// fmt.Printf("Type of c1: %T", c1)
	// fmt.Println()
	// fmt.Printf("Type of r1: %T", r1)

	// fmt.Println()
	// fmt.Println("Circle Area: ", c1.area())
	// fmt.Println("Circle Perimeter: ", c1.perimeter())

	// fmt.Println()
	// fmt.Println("Rectangle Area: ", r1.area())
	// fmt.Println("Rectangle Perimater: ", r1.perimeter())

	// c1.volume() tidak akan bisa karena telah dibuat dengan tipe shape

	// untuk penanggulan nya
	c1.(circle).volume()

	value, oke := c1.(circle)

	if oke {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f", value.volume())
	}

	// print("Rectangle", r1)
	// print("Circle", c1)

	// helpers.Reflection()
	// helpers.Goroutines()
	helpers.WaitGroup()
	fmt.Println()
}
