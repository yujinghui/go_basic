package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name string
    Age  int
    Sex  string
}

func (u User) SayHello() {
    fmt.Println("Hello there, I am %s ...", u.Name)
}

func main() {
    user := User{
        "Tony",
        30,
        "male"}
    info(user)
}

func info(o interface{}) {
    t := reflect.TypeOf(o)
    fmt.Printf("Type is: %s \n", t.Name())

    fmt.Printf("Fields are: \n")

    // find out fileds
    v := reflect.ValueOf(o)

    for i := 0; i < t.NumField(); i++ {
        f := t.Field(i)
        val := v.Field(i).Interface()
        fmt.Printf("%6s: %v=%v\n", f.Name, f.Type, val)
    }

    // find out methods
    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)
        fmt.Printf("%6s: %v\n", m.Name, m.Type)
    }
}

