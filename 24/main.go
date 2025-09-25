package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64 // private field
	y float64 // private field
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	a := NewPoint(2, 3)
	b := NewPoint(5, 9)
	distance := b.Distance(a)
	fmt.Println(distance)
}
