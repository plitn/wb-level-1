package task24

import (
	"fmt"
	"math"
)

func RunTask24() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	fmt.Println(findDistance(*p1, *p2))
}

func findDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.GetX()-p2.GetX()), 2) +
		math.Pow(float64(p1.GetY()-p2.GetY()), 2))
}

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetY() int {
	return p.y
}

func (p *Point) GetX() int {
	return p.x
}
