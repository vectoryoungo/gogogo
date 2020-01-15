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

	//Testcprint()
	//Reflecttest()
	//Showmsg()

	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}

	var mySlice []int = myArray[:5]

	fmt.Println("Element of myArray: ")

	for _, v := range myArray {
		fmt.Print(v," ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v," ")
	}
	fmt.Println()
}
