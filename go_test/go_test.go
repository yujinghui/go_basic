package testing

import (
	"fmt"
	"testing"
)

/**
* Go 要求所有的测试都在以 _test.go 结尾的文件中
8
*/

/*
* 所有的测试都应该以 func TestXxx(*testing.T) 的格式来编写。
*/
func TestHello(t *testing.T){
	fmt.Println("hell world")
	t.Errorf("can not believe it")
}