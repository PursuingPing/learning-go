package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func main() {
	var w Wheel
	var w2 Wheel2
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w2.X = 8
	w2.Y = 8
	w2.Radius = 6
	w2.Spokes = 20
	fmt.Printf("%#v\n", w)
	fmt.Printf("%#v\n", w2)
}
