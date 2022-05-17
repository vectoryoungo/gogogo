package main

import "fmt"
/*
#include <stdlib.h>
 */
import "C"

func showMsg(string2 string)  {
	fmt.Print(string2)
}
func Showmsg() {
	fmt.Println("Hello, 世界")
	showMsg("FREE OF US")
}

func Random() int{
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func maintestc() {
	Seed(100)
	fmt.Println("Random ",Random())
}

