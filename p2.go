package main

import "fmt"

func main() {
	cnp := make(chan func(), 10) // function channel with a buffer size of 10
	for i := 0; i < 4; i++ {
		go func() {
			fmt.Println("inside anon goroutine") // Added for understanding
			for f := range cnp {
				f()
				fmt.Println("inside recuring") // Added for understanding
			}
		}()
	}
	cnp <- func() {
		fmt.Println("HERE1") // Passing HERE1 to CNP
	}
	//<-cnp // Adding to wait for goroutine to finish
	<-cnp
	fmt.Println("Hello")

}

//
