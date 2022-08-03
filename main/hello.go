package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/cskr/pubsub"
	"io"
	"net/http"
	"strconv"
	"unsafe"
)

type T struct {
	A int
	B string
}

type BlockInfo struct {
	BlockHeight uint16
	BlockTx		string
}

type TxEvents struct {
	TxHash string
	Address string
}

func helloHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Hello,People Are you ok")
}

func maisn()  {
	//http.HandleFunc("/hello",helloHandler)
	//err := http.ListenAndServe(":8080",nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe:",err.Error())
	//}
	//Solutionss(10011);
	//test(100,1);
	//a := []int{1,2,3,3};
	//b := []int{2,3,1,4};
	//var result = Solution(a,b,4);
	//fmt.Print(result)
	//multiThreadRead();
	//Example();
	//fmt.Printf("％d",math.MaxUint64);
	//t := T{203,"mh203"}
	//s := reflect.ValueOf(&t).Elem()
	//typeOfT := s.Type()
	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
	//}
	//data: [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 5 245 225 0]
	byteArray := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,245,225,0};
	str1 := binary.BigEndian.Uint64(byteArray);
	fmt.Println(str1)
	var ret uint64
	buf := bytes.NewBuffer(byteArray)
	binary.Read(buf, binary.BigEndian, &ret)
	fmt.Println(ret)
	// integer for convert
	num := int64(1354321354812)
	fmt.Println("Original number:", num)

	// integer to byte array
	byteArr := IntToByteArray(num)
	fmt.Println("Byte Array", byteArr)

	// byte array to integer again
	numAgain := ByteArrayToInt(byteArr)
	fmt.Println("Converted number:", numAgain)
	encodedString := hex.EncodeToString(byteArray)

	fmt.Println("Encoded Hex String: ", encodedString)
	num, err := strconv.ParseInt(encodedString, 16, 64)

	if err != nil {

		panic(err)

	}

	fmt.Println("decimal num: ", num)
}

func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0 ; i < size ; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func ByteArrayToInt(arr []byte) int64{
	val := int64(0)
	size := len(arr)
	for i := 0 ; i < size ; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}

func Solutionss(N int) {
	var enable_print int;
	for N > 0 {
		enable_print = N % 10;     
		if enable_print == 0 && N % 10 != 0 {
			enable_print = 1;
		} else if enable_print == 1 {
			fmt.Print(N % 10);
		} else {
			fmt.Print(N % 10);
		}
		N = N / 10;
	}
}           

func test(N int,site int) {
	     fmt.Print(N/10^(site-1)%10)
}

func Solutionll(A []int,B []int,N int) int{
		edgeCount := make([]int,N);
		var m = len(A);
		var maxRank = -2147483648;

		for i:=0; i< m; i++ {
			edgeCount[A[i] - 1]++;
			edgeCount[B[i] - 1]++;
		}

		for i:=0;i<m;i++ {
		var rank = edgeCount[A[i]-1] + edgeCount[B[i]-1] - 1;
			if rank > maxRank {
				 maxRank = rank;
			}
		}
		return maxRank;
}

func multiThreadRead() {
	from := 600;
	to := 651;
	batchSize := 5;
	segament := (to-from)/batchSize;
	for i :=0; i< batchSize; i++ {
		formheight := from+segament*i;
		toheight := formheight+segament;
		go func() {
		}()
		fmt.Print("i from" ,i,formheight);
		fmt.Print("i to ",i, toheight)
		fmt.Println("");
	}
}

const topic = "topic"
const topic1="topic1"
func Example() {
	ps := pubsub.New(1)
	//pss := pubsub.New(1)
	ch := ps.Sub(topic)
	//cha := pss.Sub(topic1) BlockInfo{2322, "0x2131231231dad31231"}
	block := &BlockInfo{2322,"0x1213123123123"}
	//block.BlockHeight = 2322
	//block.BlockTx = "0x2131231231dad31231";
	//txEvent := &TxEvents{"0x1213213123", "asdfasdfsfasfa"}
	go publish(ps, block);

	//var bb = BlockInfo{998,"tssdfsdf"}
	//fmt.Println(bb)
	select {
	case msg, ok := <-ch:
		if ok {
			blockI := msg.(*BlockInfo);
			blockI.BlockHeight = 112;
			blockI.BlockTx = "txsdsfsd";
			fmt.Println(blockI.BlockTx)
			fmt.Println(blockI.BlockHeight)
		}
	}


	//select {
	//case msg,ok := <-ch:
	//	if ok {
	//		b,error := msg.(BlockInfo)
	//		if error {
	//			fmt.Println("error happen ",error)
	//		}
	//		//b := json.Unmarshal(reflect.ValueOf(msg).Elem().Bytes(),msg)
	//		//fmt.Println(b.Error())
	//		fmt.Println(b)
	//		fmt.Println(reflect.ValueOf(msg).Elem())
	//		//b := msg.(BlockInfo)
	//		//b.BlockTx
	//		//b.BlockHeight
	//		//blocks := msg.(BlockInfo);
	//		//fmt.Println(b.BlockTx)
	//		//fmt.Println(b.BlockHeight)
	//	} else {
	//		fmt.Println("wrong...")
	//	}
	//}


	//if msg, ok := <-ch; ok {
	//	fmt.Println("Received,times. ", msg)
	//	go publish(pss, txEvent)
	//	if message, ok := <-cha; ok {
	//		fmt.Println("message ", message)
	//	}
	//}

	//for i := 1; ; i++ {
	//	if i == 5 {
	//		// See the documentation of Unsub for why it is called in a new
	//		// goroutine.
	//		go ps.Unsub(ch, "topic")
	//	}
	//
	//	if msg, ok := <-ch; ok {
	//		blocks := msg.(BlockInfo);
	//		fmt.Println(blocks.BlockTx)
	//		fmt.Println(blocks.BlockHeight)
	//		fmt.Println("Received,times. ", msg, i)
	//	} else {
	//		break
	//	}
		//}
		// Output:
		// Received message, 1 times.
		// Received message, 2 times.
		// Received message, 3 times.
		// Received message, 4 times.
		// Received message, 5 times.
		// Received message, 6 times.
	//}
}

func publish(ps *pubsub.PubSub,block interface{}) {
	for {
		ps.Pub(block, topic)
	}
}