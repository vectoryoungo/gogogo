package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float64
}

/*
 * {"Title":"Go语言编程","Authors":["xushiwei","HughLv"],"Publisher":"ituring.com.cn","IsPublished":true,"Price":9.99}
 */
func main()  {
	gobook := Book{"Go语言编程",[]string{"xushiwei","HughLv"},"ituring.com.cn",true,9.99}
	b,err := json.Marshal(gobook)
	if err != nil {
		fmt.Println("error occur when marshal gobook")
	}
	fmt.Println(string(b))
	var book Book
	errs := json.Unmarshal(b,&book);
	if errs !=nil {
		fmt.Println("error when unmarshal")
	}
	fmt.Println("unmarshal result ")
	fmt.Println(book)
}