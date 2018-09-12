package main

import (
	"io"
	"strings"
	"fmt"
)

func main() {
	fmt.Println("start... ")
	r := strings.NewReader("hello world!")
	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v", n)
		fmt.Printf("%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
