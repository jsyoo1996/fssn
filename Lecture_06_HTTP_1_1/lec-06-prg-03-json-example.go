package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type members struct {
	Name           string   `json:"name"`
	Age            int64    `json:"age"`
	SecretIdentity string   `json:"string"`
	Powers         []string `json:"powers"`
}

type SuperHeroes struct {
	SquadName  string    `json:"Super hero squad"`
	HomeTown   string    `json:"homeTown"`
	Formed     int       `json:"formed"`
	SecretBase string    `json:"secretBase"`
	Active     bool      `json:"active"`
	Members    []members `json:"members"`
}

func main() {
	b, err := ioutil.ReadFile("lec-06-prg-03-json-example.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("errrorroror")
		return
	}
	data := SuperHeroes{}
	_ = json.Unmarshal([]byte(b), &data)

	fmt.Println(data.HomeTown)
	fmt.Println(data.Active)
	fmt.Println(data.Members[1].Powers[2])

}
