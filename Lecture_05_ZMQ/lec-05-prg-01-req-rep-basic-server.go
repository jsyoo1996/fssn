package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	for {
		// Wait for next request from client
		msg, _ := s.Recv(0)
                fmt.Printf("Received reply [ %s ]\n", msg)
		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		s.Send("World", 0)
	}
}

// package main

// import (
// 	"fmt"
// 	zmq "github.com/alecthomas/gozmq"
// 	"time"

// )

// func main() {
// 	context, _ := zmq.NewContext()
// 	socket, _ := context.NewSocket(zmq.REP)
// 	defer context.Close()
// 	defer socket.Close()
// 	socket.Bind("tcp://*5555")

// 	for {
// 		msg, _ := socket.Recv(0)
// 		println("Received ", string(msg))

// 		time.Sleep(time.Second)

// 		reply := fmt.Sprintf("World")
// 		socket.Send([]byte(reply), 0)
// 	}
// }
