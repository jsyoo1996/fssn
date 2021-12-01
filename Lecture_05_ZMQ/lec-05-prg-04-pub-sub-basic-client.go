package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()
	s, _ := zctx.NewSocket(zmq.SUB)
	defer s.Close()


	var temps []string
	var err error
	var temp int64
	total_temp := 0
	filter := "59937"

	if len(os.Args) > 1 {
		filter = string(os.Args[1])
	}

	fmt.Printf("Collecting updates from weather server for %s...\n", filter)
	s.SetSubscribe(filter)
	s.Connect("tcp://localhost:5556")

	for i := 0; i < 101; i++ {
		datapt, _ := s.Recv(0)
		temps = strings.Split(string(datapt), " ")
		temp, err = strconv.ParseInt(temps[1], 10, 64)

		if err == nil {
			total_temp += int(temp)
		}
		fmt.Printf("Receive temperature for zipcode %s was %dF \n\n", filter, temp)

	}

	fmt.Printf("Average temperature for zipcode %s was %dF \n\n", filter, total_temp/100)
}
