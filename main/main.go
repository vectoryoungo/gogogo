package main

import (
	"fmt"
	"log"
)

func pointereturn() *int {
	a := 0x100
	return &a
}

func pointerCopy(X *int) {
	fmt.Println("pointer: %p,target: %v\n", &X, X)
}

func secondPointer(p **int) {
	X := 100
	*p = &X
}

func variableTest(a ...int) {
	fmt.Println(a)
}

func copyArray(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

/**
 *
 */
func main() {
	var a *int = pointereturn()
	println(a, *a)

	aa := 0x100
	p := &aa
	fmt.Println("pointer: %p,target: %v\n", &p, p)
	pointerCopy(p)

	var pp *int
	secondPointer(&pp)
	println(*pp)
	//将切片作为变参时，必须进行展开操作，如果是数组，先将其转换为切片。
	ele := [3]int{7, 8, 9}
	variableTest(ele[:]...) //转换为slice后展开

	//由于变参是切片，那么参数复制的仅是切片自身，并不包括底层数组，因此可以修改原数据。
	//如果需要可用内置函数copy复制底层数据.
	arraCopy := []int{100, 200, 300}
	copyArray(arraCopy...)
	fmt.Println(arraCopy)

	defer func() {
		if err := recover(); err != nil {
			print("enter here...")
			log.Fatalln(err)
		}
	}()

	panic("The program is dead")
	println("try to exit.") // unreached statement here.
}
