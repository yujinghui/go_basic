package main

import (
	"fmt"
	"math"
)

type I interface {
	Print()
}

type People struct {
	name string
}

func (p *People) Print() {
	fmt.Println(p.Print)
}

type MyFloat1 float64

func (f MyFloat1) Print() {
	fmt.Println(f)
}

func describe(i I) {

    // %v represents the value
    // %T represents the type of the value
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
func main() {
	var i I
	i = &People{"yujinghui"}
	describe(i)
	i.Print()

	i = MyFloat1(math.Pi)
	describe(i)
	i.Print()
}
*/