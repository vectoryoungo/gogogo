package main

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
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

func deferNest() {
	defer println("one defer first appearance")
	defer println("two defer second appearance")
	panic("The program exception ...")
}

func pointerCopyPointer(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func recoverProtectCode(x, y int) {
	z := 0
	func() {
		defer func() {
			if recover() != nil {
				println("error occurred then run")
				//z = 0
			}
		}()
		println("code enter here")
		println("x is,y is,", x, y)
		z = x / y
	}()
	println("x/y=", z)
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

	//defer func() {
	//	if err := recover(); err != nil {
	//		print("enter here...")
	//		log.Fatalln(err)
	//	}
	//}()
	//
	//panic("The program is dead")
	//println("try to exit.") // unreached statement here.

	//defer func() { log.Println(recover()) }()
	//deferNest()

	//defer func() {
	//	for {
	//		if err := recover(); err != nil {
	//			println("error in here..")
	//			log.Println(err)
	//		} else {
	//			log.Fatalln("fatal")
	//		}
	//
	//	}
	//}()
	//
	//defer func() {
	//	panic("you are dead")
	//}()
	//
	//panic("program is dead")
	//recoverProtectCode(5, 0)
	//以切片语法（起始和结束索引号）返回子串时，其内部依旧指向原字节数组。
	//雨痕. Go语言学习笔记 (Chinese Edition) (Kindle Locations 1307-1308). 电子工业出版社. Kindle Edition.
	strings := "abcdefg"
	string1 := strings[:3]  //从头开始，仅指定结束索引位置
	string2 := strings[1:4] //指定开始和结束为止，返回[start,end)
	string3 := strings[2:]  //指定开始位置，返回后面全部内容
	println(string1, string2, string3)

	sstring := "hello canada!"
	pointerCopyPointer("s: %x\n", &sstring)
	bs := []byte(sstring)
	//s2 := String(bs)
	pointerCopyPointer("string to []byte,bs: %x\n", &bs)

	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := d[3:7]
	slice2 := slice1[1:3]

	for i := range slice2 {
		slice2[i] += 100
	}

	fmt.Println(d)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//implement a stack with slice
	stack := make([]int, 0, 5)
	println("stack ", len(stack))
	//enter stack
	push := func(x int) error {
		n := len(stack)
		println("push n is ", n)
		if n == cap(stack) {
			return errors.New("stack is full")
		}

		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	//pop stack
	pop := func() (int, error) {
		n := len(stack)
		println("pop n is ", n)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}

	//enter stack test
	for i := 0; i < 7; i++ {
		fmt.Printf("push %d: %v,%v\n", i, push(i), stack)
	}

	//pop stack test
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop: %d,%v,%v\n", x, err, stack)
	}

	//append data to slice tail and return it
	sliceAppend := make([]int, 0, 5)
	sliceAppend1 := append(sliceAppend, 10)
	sliceAppend2 := append(sliceAppend1, 20, 30)
	fmt.Println(sliceAppend, len(sliceAppend), cap(sliceAppend))
	fmt.Println(sliceAppend1, len(sliceAppend1), cap(sliceAppend1))
	fmt.Println(sliceAppend2, len(sliceAppend2), cap(sliceAppend2))
	//the copy of two different slice.
	sliceCopy := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceCopyOne := sliceCopy[5:8]
	fmt.Println(sliceCopyOne)  //[5 6 7]
	fmt.Println(sliceCopy[4:]) //[4 5 6 7 8 9]
	newSlice := copy(sliceCopy[4:], sliceCopyOne)
	fmt.Println(newSlice)
	fmt.Println(newSlice, sliceCopy) //3 [0 1 2 3 5 6 7 7 8 9] why no 4 and why 7 is repeated
	sliceCopyTwo := make([]int, 6)
	newCopy := copy(sliceCopyTwo, sliceCopy)
	fmt.Println(newCopy, sliceCopyTwo) //6 [0 1 2 3 5 6] why 4 is not contained in here
	//***************************
	slicecc := []int{1, 2, 3, 4, 5}
	slicecc2 := []int{5, 4, 3}
	copy(slicecc2, slicecc)
	fmt.Println("slicecc", slicecc)
	fmt.Println("slicecc2", slicecc2)

	m := make(map[string]int)
	m["canada"] = 666
	m["amerca"] = 888
	m2 := map[int]struct {
		x int
	}{
		1: {x: 100},
		2: {x: 200},
	}
	fmt.Println(m, m2)

	var map1 map[string]int
	map2 := map[string]int{}
	println(map1)
	println(map1 == nil)
	println(map2)
	println(map2 == nil)
	for s := range map2 {
		println("map2 content ", s)
	}
}
