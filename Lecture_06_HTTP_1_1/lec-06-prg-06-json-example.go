package main

import (
	"encoding/json"
	"fmt"
	"strings"
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

	b, err := json.Marshal(Heroes)
	if err != nil {
		fmt.Println(err)
	}
	var data SuperHeroes
	d := json.NewDecoder(strings.NewReader(string(b)))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)

}
