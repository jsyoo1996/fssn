package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	// "strings"
	// "os"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	ctx, _ := zmq.NewContext()
	subscriber, _ := ctx.NewSocket(zmq.SUB)
	if err := subscriber.Connect("tcp://localhost:5557"); err != nil {
		panic(err)
	}
	filter := ""

	// if len(os.Args) > 1 {
	// 	filter = string(os.Args[1])
	// }

	subscriber.SetRcvtimeo(1000)
	defer subscriber.Close()
	subscriber.SetSubscribe(filter)
	publisher, _ := ctx.NewSocket(zmq.PUSH)
	publisher.Connect("tcp://localhost:5558")
	rand.Seed(time.Now().UnixNano())

	// fmt.Printf(filter)
	
	for {
		if datapt, err3 := subscriber.Recv(0); err3 !=nil {
			if datapt == "" {
				rand := rand.Intn(100) + 1
				if rand < 10 {
					fmt.Printf("I: sending message %d\n", rand)
					if _, err1 := publisher.Send(strconv.Itoa(rand), 0); err1 != nil{
						panic(err1)
					}
					// time.Sleep(500 * time.Millisecond)
				}
			} else {
				fmt.Printf("I: received message %s\n", datapt)
			}
		} else{
			if datapt == "" {
				// fmt.Printf("Empty message\n")
				// rand := rand.Intn(100) + 1
				// if rand < 10 {
				// 	fmt.Printf("I: sending message %d\n", rand)
				// 	if _, err1 := publisher.Send(strconv.Itoa(rand), 0); err1 != nil{
				// 		panic(err1)
				// 	}
				// 	time.Sleep(500 * time.Millisecond)
				// }
			} else {
				fmt.Printf("I: received message %s\n", datapt)
			}
		}
		}

}

