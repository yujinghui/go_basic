package main

import (
	"fmt"
)

func sum(arr []int, channel chan int) {
	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	channel <- sum
}

/*
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 78, 2}
	c := make(chan int, 16)
	fmt.Println(cap(c))
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
*/