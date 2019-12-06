package main

import (
	"fmt"
	"time"
)

type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}

func main() {
	cooked := make(chan *Cake)
	iced := make(chan *Cake)

	go baker(cooked)
	go icer(iced, cooked)

	for cook := range iced {
		fmt.Println("真香", *cook)
		time.Sleep(1 * time.Second)
	}
}
