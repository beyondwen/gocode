package main

import (
	"fmt"
	"math"
)

type Move struct {
	name  string
	price float64
}

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func (m *Move) summary() string {
	return "绝地求生"
}

func main() {
	//m := Move{name: "wenhao", price: 12.5}
	//
	//fmt.Println(m.summary())

	s := Sphere{Radius: 5}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}
