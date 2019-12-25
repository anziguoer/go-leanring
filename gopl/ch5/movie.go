package ch5

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title:"Casablanca", Year:1942, Color:false, Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title:"Cool Hand Luke", Year:1976, Color:true, Actors:[]string{"Steve McQueen", "jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	//data, err := json.MarshalIndent(movies, "", "  ")
	//if err != nil {
	//	log.Fatalf("Json MarshalIndent failed: %s", err)
	//}
	//fmt.Printf("%s\n", data)

	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshal failed:%s", err)
	}

	fmt.Println(titles)
}
