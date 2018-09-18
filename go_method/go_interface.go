package  main

import (
	"fmt"
)


type Showable interface {
    showoff()
    showoff1()
}


type Person struct {
    name string
    age int32
}

func (p *Person) showoff() {
    fmt.Printf( "%v-%v" , p.name,  p.age)
}
func main(){
    var p  Showable = Person{"yujinghui", 18}
    p.showoff()
}
