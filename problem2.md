# Problem Statement 2

Explain what the following code is attempting to do? You can explain by: Giving
use-cases of what this construct/pattern could be used for?

```problem2.go
package main

import "fmt"

func main() {
	cnp := make(chan func(), 10)
	for i := 0; i < 4; i++ {
		go func() {
			for f := range cnp {
				f()
			}
		}()
	}
	cnp <- func() {
		fmt.Println("HERE1")
	}
	fmt.Println("Hello")
}

```
- Inside the main function, this creates a channel of type func  with a buffer of 10
- Inside the for loop we declare an anonymous go routine which calls the function `f()` that is supposed to pass **HERE1**
- For every iteration of the loop the program spits out another go routine which calls `f()` passing the value **HERE1** to the channel.
- Most of the time Hello is the output as main routine would finish before other go routines 


## Rewriting the code a bit

``` problem2.go
package main

import "fmt"

func main() {
	cnp := make(chan func(), 10) // function channel with a buffer size of 10
	fmt.Println("Inside main")
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
	fmt.Println("Hello")

}
```
* Running this gave

```
Inside main
inside anon goroutine
HERE1
inside recuring
inside anon goroutine
inside anon goroutine
Hello
```

- Inside the go routine when the `f()` is called **HERE1** is supposed to be passed to channel
- When multiples go routines were created inside the loop , main routine supposedly finished its execution before other routines could finish.
- Running this repeatedly might give different results. 


## Goroutines

- This is mainly used to boost concurrency and performance. Whenever "go" keyword is encountered, the go program spawns another concurrent process.
- This process runs concurrently while other code is executed.


Lets say I have 5 links and I want to check HTTP status of each link. Assuming each link is a website 

`link:=[]string{link1,link2,link3......linkn}`

I would make a request to the page and get HTTP status for each link. The code will attempt to get the status for link1, wait for reply then move to successive links

- Instead it is better to add a goroutine and pass the url go routine and run them concurrently
- It is important to block the code until all the the goroutines are finished. Using channels help.

