package main

import (
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()
	s, _ := zctx.NewSocket(zmq.PUB)
	defer s.Close()
	s.Bind("tcp://*:5556")
	s.Bind("ipc://weather.ipc")

	rand.Seed(time.Now().UnixNano())

	for {
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity)

		// Send reply back to client
		s.SendBytes([]byte(msg), 0)
	}
}
