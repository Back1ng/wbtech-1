package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) GetX() float64 {
	return p.x
}

func (p Point) GetY() float64 {
	return p.y
}

func getDistanceBetweenPoints(pointA, pointB Point) float64 {
	// в данном случае находясь в одном пакете, можно обращаться напрямую - pointA.x, pointA.y
	return math.Sqrt(math.Pow(pointB.GetX()-pointA.GetX(), 2) + math.Pow(pointB.GetY()-pointA.GetY(), 2))
}

func main() {
	a := NewPoint(2, 1)
	b := NewPoint(3, 2)

	fmt.Println(getDistanceBetweenPoints(a, b))
}
