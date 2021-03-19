package main

import (
	"fmt"
	"time"
)

// ping pong using concurrency
// by gorutines and channels

// pinger
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// ponger
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// start the gorutine
	ping <- 1

	for {
		// block the main threads until an interrupt
		time.Sleep(time.Second)
	}
	// option to for{}
	// select {}
}
