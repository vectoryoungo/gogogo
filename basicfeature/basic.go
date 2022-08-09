package basicfeature

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

// will auto increment with _
//const (
//	_ int = iota
//	Foo
//	Bar
//	_
//	_
//	Bin
//	Baz
//)

//will not auto increment with blank line.
const (
	_ int = iota
	Foo
	Bar

	Bin
	Baz
)

//bitmasks operate with iota
const (
	read = 1 << iota
	write
	remove

	admin = read | write | remove
)

//calculate memory size
//const (
//	KB = 1024       // binary 00000000000000000000010000000000
//	MB = 1048576    // binary 00000000000100000000000000000000
//	GB = 1073741824 // binary 01000000000000000000000000000000
//)

//use both shift operator and a multiplier
const (
	_  = 1 << (iota * 10)
	KB // decimal:       1024 -> binary 00000000000000000000010000000000

	MB // decimal:    1048576 -> binary 00000000000100000000000000000000

	GB // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

/*
* You can also pair up constants on the same line with iota .
  As well as you can use an underscore to skip a value within those pairs.
  Here is an example of going crazy with iota :
*/
const (
	tomato, apple int = iota + 1, iota + 2
	orange, chevy
	ford, _
)

//可以为当前包内的任意类型定义方法
//雨痕. Go语言学习笔记 (Chinese Edition) (Kindle Location 201). 电子工业出版社. Kindle Edition.
type X int

func (x *X) inc() {
	*x++
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero...")
	}

	return a / b, nil
}

func Basic() {
	a, b := 10, 2
	c, err := div(a, b)
	fmt.Println(c, err)

	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	z, ok := m["a"]
	fmt.Println("x is ", x, ok)
	fmt.Println("z is ", z, ok)

	delete(m, "a")

	y, ok := m["b"]

	fmt.Println("y is ", y, ok)

	var man manager
	man.name = "tom"
	man.age = 35
	man.title = "senior manager"
	fmt.Println("manager description:", man)

	var xx X
	xx.inc()
	fmt.Println(xx)

	const babysit = "this is babysit"
	const endingSymbol = "this is ending symbol"

	fmt.Println("The value of Red is %v\n", Red)
	fmt.Println("The value of Orange is %v\n", Orange)
	fmt.Println("The value of Yellow is %v\n", Yellow)
	fmt.Println("The value of Green is %v\n", Green)
	fmt.Println("The value of Blue is %v\n", Blue)
	fmt.Println("The value of Indigo is %v\n", Indigo)
	fmt.Println("The value of Violet is %v\n", Violet)

	fmt.Println("The value of Foo is %v\n", Foo)
	fmt.Println("The value of Bar is %v\n", Bar)
	fmt.Println("The value of Bin is %v\n", Bin)
	fmt.Println("The value of Baz is %v\n", Baz)

	/**
	 * read   = 1 << iota // 00000001 = 1
	 * write              // 00000010 = 2
	 * remove             // 00000100 = 4
	 *
	 */
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)

	fmt.Printf("KB =  %v\n", KB)
	fmt.Printf("MB =  %v\n", MB)
	fmt.Printf("GB =  %v\n", GB)

	fmt.Printf("tomato =  %v\n", tomato)
	fmt.Printf("apple =  %v\n", apple)
	fmt.Printf("orange =  %v\n", orange)
	fmt.Printf("chevy =  %v\n", chevy)
	fmt.Printf("ford =  %v\n", ford)
}

func Hello() {
	println("hello function test")
}

func Exec(f func()) {
	f()
}
