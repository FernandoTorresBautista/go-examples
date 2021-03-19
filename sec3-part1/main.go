package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//21 Go rutines
	// goroutines()

	// 22 go routines continued
	goroutines2()
}

func goroutines2() {
	// wait groups
	wg := &sync.WaitGroup{}

	// mutexes ( critical section)
	mut := &sync.Mutex{}
	mut.Lock()

	// add, done and wait functions
	wg.Add(2)
	// lambda functions - anoninous functions
	func() {
		fmt.Println("hello ")
		wg.Done()
	}() // calling the function with ()

	func() {
		fmt.Println("world")
		wg.Done()
	}()
	//
	mut.Unlock()

	fmt.Println("Fin print before end waits")
	wg.Wait() // wait after 2 dones
	fmt.Println("Fin 2")
	//// undefined cases
	//select {}

}

// part 1 go routines

func heavy() {
	count := 0
	for {
		if count > 5 {
			break
		}
		time.Sleep(time.Second * 1)
		fmt.Println("heavy")
		count++
	}
	//time.Sleep(time.Second * 5)
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("super heavy")
	}
}

func goroutines() {
	//heavy()
	go heavy() // run the function in the background and don't stop anything below this line
	go superHeavy()
	fmt.Println("end")

	// if don't put this no execite nothing in go back, acause main func end well
	//time.Sleep(time.Second * 5)

	// with this the go routine continue before stop
	select {}
}
