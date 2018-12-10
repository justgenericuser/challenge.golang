package main

import (
	"log"
	"sync"
	"fmt"
	"time"
	"math/rand"
)
func main() {
	problem2()
}



func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	// create a new instance of a sync.WaitGroup, named wg
	var wg sync.WaitGroup


	//creating the ticker with 1sec as time interval
	ticker := time.NewTicker(1000 * time.Millisecond)

	for inx := 0; inx < 10; inx++ {
		// add another go routine to WaitGroup in order to wait for it to finish executing before method problem1 finishes its execution
		wg.Add(1)
		go printRandom2(inx, &wg, *ticker)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	//time.Sleep(5 * time.Second)
	// wait for all the go routines to finish
	wg.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, wg *sync.WaitGroup, ticker time.Ticker) {
	// schedule the call to Done to tell func problem2 we are done
	// defer is used to call the wg.Done() method at the end of the enclosing function, i.e., printRandom2
	defer wg.Done()

	for inx := 0; inx < 10; inx++ {

		for t := range ticker.C {
			fmt.Println("problem2: Tick at", t, "slot =", slot, "count =", inx, "rand =",rand.Int())
			//log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
			break
		}
	}
}
