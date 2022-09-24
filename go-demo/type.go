package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(len("Go语言")) // 8
	fmt.Println(len([]rune("Go语言")))
	l := list.List{}
	list.Element{}
	fmt.Println(rune('x'))
	fmt.Println(string(20320))
	temp := []rune{20320, 22909, 32, 19990, 30028}
	fmt.Println(string(temp))

	var str string = "hello world"
	fmt.Println("byte=", []byte(str))
	fmt.Println("byte=", []rune(str))
	fmt.Println(str[:2])
	fmt.Println(string([]rune(str)[:2]))

}
