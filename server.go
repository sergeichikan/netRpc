package main

import (
	"fmt"
	"net"
	"net/rpc"
	"netRpc/UnitTypes"
)

type Reply struct {
	Data []UnitTypes.Message
}

type Listener struct {}
func (l *Listener) GetLine(data []UnitTypes.Message, reply *Reply) error {
	fmt.Printf("Receive: %v\n", len(data))
	for i := 0; i < len(data); i++ {
		data[i].Body = data[i].Body + "change"
	}
	*reply = Reply{data}
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		fmt.Println(err)
		return
	}
	listener := new(Listener)
	err = rpc.Register(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.Accept(inbound)
}
