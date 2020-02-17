package main

import (
	"fmt"
)
//most of the programming languages use os threads means a individual function call stack dedicated to execution what ever code handled by thread
//abstraction of a thread called go routine
//inside of go run time we have schedular that gonna map go routines on to the operating system threads for periods of time
//schedular within take turns with every cpu thread that is available and assign different go routines at certain amoount processing time on those threads
//but we dont have to interact with low level threads directly  we interact with high level go routines


//advantage since we have the abstraction go routines can starts with very very small stack basis, they can be reallocated very very quickly
//they are very cheap to create and destroy
func main() {
	go sayHello()  // we created go routine

}//when we run the it doesn't print the because the main function actually executing go routine itself
// in line 16 we have span another go routine the application exists as soon as main is done

func sayHello() {
	fmt.Println("Hello")
}