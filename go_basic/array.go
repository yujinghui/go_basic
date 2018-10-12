package main

import "fmt"

func main() {

	// simple declare
	var arr [10]int
	var pstr [10]*float64
	a := [3]int{1, 2, 3}
	b := [5]int{11, 2, 3}
	c := [...]int{1, 2, 3}
	d := [...]int{8: 18}
	fmt.Println(arr, pstr, a, b, c, d)
	// element visit
	intArray := []int{1, 2, 3, 4, 5}
	modify_array(intArray)
	fmt.Println(intArray)
}

/**
error: 
    intArray := [...]int{1, 2, 3, 4, 5} -----------this is array
	modify_array(intArray)
correct:
    intArray := []int{1, 2, 3, 4, 5}    -----------this is slice
	modify_array(intArray)
*/

func modify_array(arr []int) {
	arr[0] = 100
	fmt.Println(arr)
}
