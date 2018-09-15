package main

import (
	"container/list"
	"fmt"
)

/**
* go has no set ,just list and map
*/
 func main(){
	 list := list.New()
	 list.PushBack("hello")
	 list.PushBack("world")
	 fmt.Println(list)
	 fmt.Println(list.Front().Value)
	 fmt.Println(list.Back().Value)
 }