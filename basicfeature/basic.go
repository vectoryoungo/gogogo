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
}
