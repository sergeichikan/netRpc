package main

import (
	"fmt"
	"net/rpc"
	"netRpc/UnitTypes"
	"time"
)

const cycleLimit = time.Millisecond * 100
const dataLength = 100_000

type Reply struct {
	Data []UnitTypes.Message
}

var Data []UnitTypes.Message

func init() {
	for i := 0; i < dataLength; i++ {
		Data = append(Data, UnitTypes.Message{"asdas", "b", time.Now().Unix(), 1.2})
	}
	fmt.Println(len(Data))
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	var reply Reply
	for {
		startTime := time.Now()
		err = client.Call("Listener.GetLine", Data, &reply)
		if err != nil {
			fmt.Println(err)
			return
		}
		Data = reply.Data
		workTime := time.Since(startTime)
		sleepTime := cycleLimit - workTime
		fmt.Println(sleepTime)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}
}
