package main

import (
	"fmt"
	"reflect"
)

type Showable interface {
    Showoff()
    Showoff1()
}


type Person struct {
    Name string
    Age int32
}

func (p Person) Showoff() {
    fmt.Printf( "show off %v-%v\n" , p.Name,  p.Age)
}

func (p Person) Showoff1() {
	fmt.Printf( "show off1 :%v-%v\n" , p.Name,  p.Age)
}


func main(){
	var a Showable = Person{"yujinghui", 18}
	a.Showoff()
	a.Showoff1()
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf(" type of a : %s \n" + t.Name())

	for i:=0; i < t.NumField(); i++ {
           f := t.Field(i)
	   val := v.Field(i).Interface()
	   fmt.Printf("%6s: %v - %v \n", f.Name, f.Type, val)
	}
	fmt.Printf("method list as follows %v\n", t.NumMethod())
	for i:=0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s - %v\n", m.Name, m.Type)
	}
}
