package main

import (
        "fmt"

        zmq "github.com/pebbe/zmq4"
)

func main() {
        zctx, _ := zmq.NewContext()

        // Socket to talk to server
        fmt.Printf("Connecting to the server...\n")
        s, _ := zctx.NewSocket(zmq.REQ)
        s.Connect("tcp://localhost:5555")

        // Do 10 requests, waiting each time for a response
        for i := 0; i < 10; i++ {
                fmt.Printf("Sending request %d...\n", i)
                s.Send("Hello", 0)

                msg, _ := s.Recv(0)
                fmt.Printf("Received reply %d [ %s ]\n", i, msg)
        }
}

// package main() {
// 	context, _ := zmq.NewContext()
// 	socket, _ := context.NewSocket(zmq.REQ)
// 	defer context.Close()
// 	defer socket.Close()

// 	fmt.Printf("Connecting to hello world server...")
// 	socket.Connect("tcp://localhost:5555")

// 	for i := 0; i< 10; i++ {
// 		msg := fmt.Sprintf("Hello %d", i)
// 		socket.Send([]byte(msg), 0)
// 		println("Sending ", msg)

// 		reply, _ := socket.Recv(0)
// 		println("Received ", string(reply))
// 	}
// }

