package main

import (
	zmq "github.com/pebbe/zmq4"
	"fmt"
)

func main() {
	ctx, _ := zmq.NewContext()
	publisher, _  := ctx.NewSocket(zmq.PUB)
	publisher.Bind("tcp://*:5557")
	collector, _ := ctx.NewSocket(zmq.PULL)
	if err := collector.Bind("tcp://*:5558"); err != nil {
		panic(err)
	}
	// publisher.Bind("ipc://number.ipc")

	for {
		if msg, err := collector.Recv(0); err !=nil{
			panic(err)
		} else{
			fmt.Printf("I: publishing update %s\n", msg)
			if _, err1 := publisher.Send(msg, 0); err1 != nil {
				panic(err1)
			}
		}
	}
}