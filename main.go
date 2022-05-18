package main

import (
	"fmt"
	// "time"
)

var finished = make(chan bool)

// Here, the value of Sleep function is zero
// So, this function return immediately.
func show(str string) {

	for x := 0; x < 4; x++ {

		// time.Sleep(2 * time.Millisecond)
		fmt.Println(str)

	}
	finished <- true
}

// Main Function
func main() {

	// Calling Goroutine
	go show("Hello")
	<-finished
	// Calling function
	show("Bye")

}
