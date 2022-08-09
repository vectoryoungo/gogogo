package main

import (
	"fmt"
	"github.com/segmentio/encoding/json"
	"gogogo/basicfeature"
	"io/ioutil"
	"log"
	"net/http"
)

type client struct {
}

var clientMap map[string]client

//func (c *client) insertLightningConnectorIntoComputer(com computer) {
//	fmt.Println("Client inserts Lightning connector into computer.")
//	com.insertIntoLightningPort()
//}

func mainner() {
	//clientMap = make(map[string] client)
	//value,ok := clientMap["client"]
	//if ok {
	//	fmt.Println("value is ",value)
	//}
	//clientMap["client"] = client{};
	//values,okk := clientMap["client"]
	//
	//if okk {
	//	fmt.Println("second value",values)
	//}
	//
	//
	//
	//
	//adidasFactory, _ := getSportsFactory("adidas")
	//nikeFactory, _ := getSportsFactory("nike")
	//
	//nikeShoe := nikeFactory.makeShoe()
	//nikeShirt := nikeFactory.makeShirt()
	//
	//adidasShoe := adidasFactory.makeShoe()
	//adidasShirt := adidasFactory.makeShirt()
	//
	//printShoeDetails(nikeShoe)
	//printShirtDetails(nikeShirt)
	//
	//printShoeDetails(adidasShoe)
	//printShirtDetails(adidasShirt)
	//
	//client := &client{}
	//mac := &mac{}
	//
	//client.insertLightningConnectorIntoComputer(mac)
	//
	//windowsMachine := &windows{}
	//windowsMachineAdapter := &windowsAdapter{
	//	windowMachine: windowsMachine,
	//}
	//
	//client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
	//client := binance.NewClient("5vWfkszWGjFZcmaOA5kelER1VAT6IEWB1cMfy8AaAiE3pyY4uR9kEHL2bRMfqpLd", "3jCRUolx0YtFqqvrhNJniPwpRY2vQooPHPOiRJ7GkF574UJCToRCgabDB3brUEeB")
	//
	//err := client.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//
	//prices, err := client.Prices()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, p := range prices {
	//	log.Printf("symbol: %s, price: %s", p.Symbol, p.Price)
	//}

	Test()
	basicfeature.Basic()
	done := make(chan bool)
	data := make(chan int)

	go basicfeature.Consumer(data, done)
	go basicfeature.Producer(data)
	<-done //阻塞，直到消费者发回结束信号 雨痕. Go语言学习笔记 (Chinese Edition) (Kindle Location 243). 电子工业出版社. Kindle Edition.
	// Create default client
	//client := binance.NewClient("API-KEY", "SECRET")
	//
	//// Send ping request
	//err := client.Ping()
	//
	//// Create client with custom request window size
	//client := binance.NewClient("API-KEY", "SECRET").ReqWindow(5000)
	//
	//// Create websocket client
	//wsClient := ws.NewClient()
	//
	//// Connect to Klines websocket
	//ws, err := wsClient.Klines("ETHBTC", binance.KlineInterval1m)
	//
	//// Read ws
	//msg, err := ws.Read()
	f := basicfeature.Hello
	basicfeature.Exec(f)
}

//func printShoeDetails(s iShoe) {
//	fmt.Printf("Logo: %s", s.getLogo())
//	fmt.Println()
//	fmt.Printf("Size: %d", s.getSize())
//	fmt.Println()
//}
//
//func printShirtDetails(s iShirt) {
//	fmt.Printf("Logo: %s", s.getLogo())
//	fmt.Println()
//	fmt.Printf("Size: %d", s.getSize())
//	fmt.Println()
//}

func Test() {
	//req, err := http.NewRequest("GET", EtherscanAPIMapping[chain], nil)
	req, err := http.NewRequest("GET", "http://api.etherscan.io/api", nil)
	if err != nil {
		fmt.Errorf("http NewRequest err")
	}

	q := req.URL.Query()
	q.Add("module", "account")
	q.Add("action", "txlistinternal")
	q.Add("address", "0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3")
	q.Add("startblock", "0")
	q.Add("endblock", "2702579")
	q.Add("page", "1")
	q.Add("offset", "10")
	q.Add("sort", "desc")
	q.Add("apikey", "8Y1TZVGAXMQ5XPZ4Y5PEXCG5N3FTBF6HBN")

	req.URL.RawQuery = q.Encode()

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
	var repResult RepResult
	json.Unmarshal(result, &repResult)
	fmt.Println("repResult is ", repResult)
	jsonResult, _ := json.Marshal(repResult)
	fmt.Println("jsonResult", jsonResult)
	defer resp.Body.Close()

	fmt.Println("res: ", resp.Body)
}

type RepResult struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}
type Result struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceID         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}
