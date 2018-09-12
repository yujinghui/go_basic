package main

import "fmt"
/*
func main() {
	var myi interface{}
	describe1(myi)
}
*/
func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
