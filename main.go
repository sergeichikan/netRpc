package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"netRpc/UnitTypes"
	"time"
)

const defaultAddress = "localhost:3000"

type Listener struct{}

func (l *Listener) Handler(data []UnitTypes.Message, reply *UnitTypes.Reply) error {
	fmt.Printf("Receive: %v\n", len(data))
	timestamp := time.Now().Unix()
	for i := range data {
		data[i].Time = timestamp
	}
	*reply = UnitTypes.Reply{data}
	return nil
}

func run(address string) {
	addy, err := net.ResolveTCPAddr("tcp", address)
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

func main() {
	address := flag.String("address", defaultAddress, "address to server")
	flag.Parse()
	fmt.Println("address", *address)
	run(*address)
}
