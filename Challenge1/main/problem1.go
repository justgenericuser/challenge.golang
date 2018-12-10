package main

import (
	"log"
	"sync"
	"math/rand"
)

func main() {
	problem1()
}

func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	// create a channel named 'print' which takes boolean messages
	printNum := make(chan bool)

	// create a channel named quit, this is used for signalling the go routines to quit / return
	quit := make(chan struct{})

	// create a new instance of a sync.WaitGroup, named wg
	var wg sync.WaitGroup

	for inx := 0; inx < 10; inx++ {

		// add another go routine to WaitGroup in order to wait for it to finish executing before method problem1 finishes its execution
		wg.Add(1)

		// call the method printRandom1 as a concurrent go routine, pass it the reference for WaitGroup, and channels 'print' and 'quit'
		go printRandom1(inx, &wg, printNum, quit)

	}

	// receive on channel 'print' 100 times
	for inx := 0; inx < 100; inx++ {
		// receive the message from channel 'print', each receive will cause one of the concurrent go routines to print a random number
		<-printNum
	}

	// 100 random numbers printed

	// close the channel 'quit', close is used to broadcast to all (10) concurrent go routines to return and quit
	close(quit)

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	// wait for all the go routines to finish
	wg.Wait()

	// close the channel 'print' after all go routines finish
	close(printNum)

	log.Printf("problem1: finised --------------------------------------------")
}

// add the following parameters to the function printRandom1
// 1. the pointer to WaitGroup instance, wg *sync.WaitGroup
// 2. the channel named print (used for message passing, boolean)
// 3. the channel named quit (used for signaling the go routines to return / quit)
func printRandom1(slot int, wg *sync.WaitGroup, printNum chan bool, quit chan struct{}) {

	//
	// Do not change 25 into 10!
	//

	// schedule the call to Done to tell func problem1 we are done
	// defer is used to call the wg.Done() method at the end of the enclosing function, i.e., printRandom1
	defer wg.Done()

	for inx := 0; inx < 25; inx++ {

		// the following select blocks until one of its cases can run, then it executes that case. it chooses one at random if multiple are ready.
		select {
		// receive from channel 'quit', it is a blocking receive, it will block until the channel is closed
		case <-quit:
			// received on quit, quitting the go routine
			return
			// send 'true' to channel 'print', it is a blocking send, it will block until we receive on channel 'print' in func problem1
		case printNum <- true:
			// print the random number
			log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		}
	}
}
