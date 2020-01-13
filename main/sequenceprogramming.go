package main

import "fmt"

func main()  {
	var i int;
	var j int;
	i = 10
	j = 11
	fmt.Println(" i is ",i)
	fmt.Println(" j is ", j)
	i,j = j,i
	fmt.Println("i is %d ",i)
	fmt.Println("j is %d ",j)
}
