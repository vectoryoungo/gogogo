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
首先，在repeatedlyLock函数被执行伊始，对互斥锁的第一次锁定操作便被进行并顺利地完成。
这由第一行和第二行打印内容可以看出。
而后，在repeatedlyLock函数中被启用的那三个Goroutine在G0的第一次“睡眠”期间开始被运行。
当相应的go函数中的对互斥锁的锁定操作被进行的时候，它们都被阻塞住了。
原因是该互斥锁已处于锁定状态了。
这就是我们在这里只看到了三个连续的Lock the lock. (G)而没有立即看到The lock is locked. (G)的原因。
随后，G0“睡醒”并解锁互斥锁。
这使得正在被阻塞的G1、G2和G3都会有机会重新锁定该互斥锁。
但是，只有一个Goroutine会成功。
成功完成锁定操作的某一个Goroutine会继续执行在该操作之后的语句。
而其他Goroutine将继续被阻塞，直到有新的机会到来。
这也就是上述打印内容中的最后三行所表达的含义。
显然，G1抢到了这次机会并成功锁定了那个互斥锁。
实际上，我们之所以能够通过使用互斥锁对共享资源的唯一性访问进行控制正是因为它的这一特性。
这有效的对竞态条件进行了消除。互斥锁的锁定操作的逆操作并不会引起任何Goroutine的阻塞。
但是，它的进行有可能引发运行时panic。更确切的讲，当我们对一个已处于解锁状态的互斥锁进行解锁操作的时候，就会已发一个运行时panic。
这种情况很可能会出现在相对复杂的流程之中——我们可能会在某个或多个分支中重复的加入针对同一个互斥锁的解锁操作。
避免这种情况发生的最简单、有效的方式依然是使用defer语句。这样更容易保证解锁操作的唯一性。
虽然互斥锁可以被直接的在多个Goroutine之间共享，但是我们还是强烈建议把对同一个互斥锁的成对的锁定和解锁操作放在同一个层次的代码块中。
例如，在同一个函数或方法中对某个互斥锁的进行锁定和解锁。
又例如，把互斥锁作为某一个结构体类型中的字段，以便在该类型的多个方法中使用它。
此外，我们还应该使代表互斥锁的变量的访问权限尽量的低。
这样才能尽量避免它在不相关的流程中被误用，从而导致程序不正确的行为。
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
func mainse() {
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
