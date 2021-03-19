package main

import (
	"fmt"
	"os"
)

func Select(c chan int, quits chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("received")
		case <-quits:
			fmt.Println("quit")
			os.Exit(0)
		}
	}
}

func main() {

	c := make(chan int)
	quits := make(chan struct{})

	go func() {
		c <- 1

		quits <- struct{}{}
	}()

	go Select(c, quits)

	select {}
	// why select{} : forma de decirle que no ha terminado
	// controlar la ejecución
}
