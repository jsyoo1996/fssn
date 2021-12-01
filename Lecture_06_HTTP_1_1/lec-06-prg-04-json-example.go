package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type members struct {
	Name           string
	Age            int64
	Secretidentity string
	Powers         []string
}

type SuperHeroes struct {
	SquadName  string
	HomeTown   string
	Formed     int
	SecretBase string
	Active     bool
	Members    []members
}

func main() {
	Heroes := `
	[{
		"squadName": "Super hero squad",
 		"homeTown": "Metro City",
 		"formed": 2016,
 		"secretBase": "Super tower",
 		"active": true,
		"members": [{
			"name": "Molecule Man",
			"age": 29,
			"secretIdentity": "Dan Jukes",
			"powers": [
				"Radiation resistance",
				"Turning tiny",
				"Radiation blast"
			]
		},
		{
			"name": "Madame Uppercut",
			"age": 39,
			"secretIdentity": "Jane Wilson",
			"powers": [
			  "Million tonne punch",
			  "Damage resistance",
			  "Superhuman reflexes"
			]
		},
		{
			"name": "Eternal Flame",
			"age": 1000000,
			"secretIdentity": "Unknown",
			"powers": [
			  "Immortality",
			  "Heat Immunity",
			  "Inferno",
			  "Teleportation",
			  "Interdimensional travel"
			]
		}
		]
	}]
	`

	var data []SuperHeroes

	_ = json.Unmarshal([]byte(Heroes), &data)

	fmt.Println(data[0].HomeTown)
	fmt.Println(data[0].Active)
	fmt.Println(data[0].Members[1].Powers[2])

	file, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile("lec-06-prg-04-json-example.json", file, 0644)
}
