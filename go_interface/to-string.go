package main

import "fmt"


type Person struct {
	Name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.age)
}
func main(){
	p := Person{"yujinghui", 18}
	fmt.Print(p)
}