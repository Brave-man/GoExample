package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url: %s\n", r.URL)
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
