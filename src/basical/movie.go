package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "CC", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "AA", Year: 1955, Color: true,
		Actors: []string{"Bogart", "Ingrid Bergman"}},
	{Title: "BB", Year: 1963, Color: true,
		Actors: []string{"Hufrt", "Ingrid Bergman"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err := json.MarshalIndent(movies, "", "     ")
	if err != nil {
		log.Fatalf("JSON marshaling failed : %s", err)
	}
	fmt.Printf("%s\n", data2)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
