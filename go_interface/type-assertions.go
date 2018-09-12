package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float32)
	fmt.Println(f, ok)

	/**
	error i does not contains float32
	*/
	// fa := i.(float32)
	// fmt.Println(fa)
}
