package main

import (
	"fmt"
)

type I interface {
	M()
	N()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Print("M:" + t.S)
}

func (t T) N() {
	fmt.Print("N:" + t.S)
}

/*
func main() {
	var i I = T{"hello"}
	i.M()
	i.N()
}
*/