package main

import (
	"fmt"
	"sync"
	"time"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func printVariable() {
	var i int
	var j int
	i = 10
	j = 11
	fmt.Println(" i is ", i)
	fmt.Println(" j is ", j)
	i, j = j, i
	fmt.Println("i is %d ", i)
	fmt.Println("j is %d ", j)
}

func printSlice() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mySlice []int = myArray[:5]

	fmt.Println("Element of myArray: ")

	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	mySlice1 := make([]int, 5, 10)

	for _, v := range mySlice1 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	mySlice3 := []int{99, 88, 77, 66, 55}

	for i := 0; i < len(mySlice3); i++ {
		fmt.Println("mySlice3[", i, "] = ", mySlice3[i])
	}

	for i, v := range mySlice3 {
		fmt.Println("mySlice3 range [", i, "] = ", v)
	}

	slice1 := []int{8, 9, 5, 7, 3}
	slice2 := []int{66, 55, 44}

	for i, v := range slice1 {
		fmt.Println("origin slice1 range [", i, "] = ", v)
	}

	for i, v := range slice2 {
		fmt.Println("origin slice2 range [", i, "] = ", v)
	}

	copy(slice2, slice1)

	for i, v := range slice2 {
		fmt.Println("copy after slice2 range [", i, "] = ", v)
	}

	copy(slice1, slice2)

	for i, v := range slice1 {
		fmt.Println("copy after slice1 range [", i, "] = ", v)
	}
}

func PrintMap() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"1", "Tim", "china chaoyang...."}
	personDB["free"] = PersonInfo{"2", "Donald Trimph", "American"}

	person, ok := personDB["free"]

	if ok {
		fmt.Println("we found us president!!!", person.Name)
	}
}

func condition(x int) int {
	if x == 5 {
		return 5
	} else {
		return x
	}
}

func switchCondition(x int) {
	switch x {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		//fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("nothing")
	}
}

func maine() {
	//Testcprint()
	//Reflecttest()
	//Showmsg()
	//printSlice()
	//result := condition(8)
	//fmt.Println(" result is ",result)
	//
	//switchCondition(2)
	//
	//sum,_ := vmath.Add(2,3)
	//println(sum)

	//
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n ", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
/**
result:
Lock the lock. (G0)
The lock is locked. (G0)
Lock the lock. (G1)Lock the lock. (G2)Lock the lock. (G3)Unlock the lock. (G0)
The lock is unlocked. (G0)
The lock is locked. (G1)
 */
func repeatedlyLock() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (G0)")
	mutex.Lock()
	fmt.Println("The lock is locked. (G0)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (G%d)", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (G0)")
	time.Sleep(time.Second)
}
func main() {
	//s1 := "100"
	//i1, err := strconv.Atoi(s1)
	//if err != nil {
	//	fmt.Println("can't convert to int")
	//} else {
	//	fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	//}
	//
	//i2 := 200
	//s2 := strconv.Itoa(i2)
	//fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
	//
	//arith := new(vmath.Arith)
	//rpc.Register(arith)
	//rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	//if e != nil {
	//	log.Fatal("listen error ", e)
	//}
	//go http.Serve(l, nil)
	//fmt.Println(" server end ")

	//os.Mkdir("test", 0777)
	//os.MkdirAll("test/test1/test2", 0777)
	//errs := os.Remove("test")
	//
	//if errs != nil {
	//
	//	fmt.Println(err)
	//}
	//
	//os.RemoveAll("test")

	repeatedlyLock()

}
