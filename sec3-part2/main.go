package main

import (
	"fmt"
	"time"
)

func main() {

	// declare and use channels

	// Unbuffered channels
	//var c chan int // empty channel print nil
	c := make(chan int)
	// send in go rutine
	go func() {
		c <- 1 // receive the datatype of the channel
	}()
	val := <-c
	fmt.Println(val)

	// <name> chan <datatype>
	fmt.Println(c) // the channel directtion memory

	go func() {
		c <- 2 // receive the datatype of the channel
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val) // the channel directtion memory

}

/*
Go Channels

-Producer <- Consumer

If u have multiple consumers
-Producer-> Channel:
	-c1
	-c2
	...
	-cn

the channel put the info from one channel to anhoter channel

*/
