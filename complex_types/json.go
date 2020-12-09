package complex_types

import (
	"encoding/json"
	"fmt"
)

func JsonEntry() {
	//makeJson()
	unmarshal()
}

func unmarshal() {
	var m Movie
	var str = "{\"title\":\"Casablanca\",\"release\":1942,\"color\":false,\"Actors\":[\"Humphrey Bogart\",\"Ingrid Bergman\"]}\n"
	json.Unmarshal([]byte(str), &m)
	fmt.Printf("%#v\n", m)
}

type Movie struct {
	Title    string `json:"title"`
	Year     int `json:"release"`
	Color    bool `json:"color"`
	Actors   []string `actors`
	duration int
}

func makeJson() {
	m := Movie{
		Title: "Casablanca",
		Year:  1942,
		Actors: []string{
			"Humphrey Bogart",
			"Ingrid Bergman",
		},
		duration: 1000,
	}
	fmt.Println(m)

	movieJson, err := json.Marshal(m)
	if err == nil {
		fmt.Printf("%s\n", movieJson)
	}

	movieJson2, err2 := json.MarshalIndent(m, "", "\t")
	if err2 == nil {
		fmt.Printf("%s\n", movieJson2)
	}
}
