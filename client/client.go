package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"netRpc/UnitTypes"
	"time"
)

const defaultCycleLimit = time.Millisecond * 100
const defaultDataLength = 100_000
const defaultAddress = "localhost:3000"

var Data []UnitTypes.Message

func initData(length int) {
	for i := 0; i < length; i++ {
		Data = append(Data, UnitTypes.Message{"name", "body", time.Now().Unix(), 1.2})
	}
	fmt.Println("length", len(Data))
}

func runClient(address string, cycleLimit time.Duration, dataLength int) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	initData(dataLength)
	var reply UnitTypes.Reply
	for {
		startTime := time.Now()
		err = client.Call("Listener.Handler", Data, &reply)
		if err != nil {
			fmt.Println(err)
			return
		}
		Data = reply.Data
		sleepTime := cycleLimit - time.Since(startTime)
		fmt.Println(sleepTime)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}

func main() {
	address := flag.String("address", defaultAddress, "address to server")
	cycleLimit := flag.Int("cycle", int(defaultCycleLimit), "time limit of iteration")
	dataLength := flag.Int("length", defaultDataLength, "data length")
	flag.Parse()
	fmt.Println("address", *address)
	fmt.Println("cycle", time.Duration(*cycleLimit))
	fmt.Println("length", *dataLength)
	runClient(*address, time.Duration(*cycleLimit), *dataLength)
}
