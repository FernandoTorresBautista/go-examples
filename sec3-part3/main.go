package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	//
	//c := make(chan int, 3) // buffer of 3
	c := make(chan *Car, 3) // buffer of 3

	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"} // use the close channel for more that 3 values only
		close(c)       //
	}()

	for i := range c {
		//fmt.Println(i)
		fmt.Println(i.Model)
	}

}
