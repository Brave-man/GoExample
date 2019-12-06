package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"` // omitempty表示这个值为空或者为零值时,不输出这个值到JSON中
	Actors []string `json:"actors"`
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Ingrid Bergman"}},
	}

	//data, err := json.Marshal(movies) // 紧凑模式
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{Title string}
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatal(err)
	}
	fmt.Println(titles)
}
