package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Abser interface {
	Abs() float64
	Abs_hha() float64
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs_hha() float64 {
	fmt.Println("v --> abs_hha")
	return 0
}
func (f MyFloat) Abs_hha() float64 {
	fmt.Println("f --> abs_hha")
	return 0
}

/*
 func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	a = &v
	fmt.Println(a.Abs())
	fmt.Println(a.Abs_hha())
}
*/
