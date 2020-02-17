package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")  
	fmt.Println("End")

	// if we defer all statements we will get last in first out mannet means " End " comes first
	// then middle and then start mean LIFO order
	// when the execution starts if any defer functions found.it calls before the main function returns
	// note defer statements only take function calls not function

	a := "Start"
	defer fmt.Println(a)
	a = "end"
	// takes the argument at the time of defer is called not at the time of function execution

	// we don't have exceptions in go ulike other languages
	// for example if want to open a file and that file does't exist that is  an actually normal response
	// for that we return error values  don't throw exceptions in go
	// however there are somethings that a go application into a situation where it can not continue

	fmt.Println("start")
	panic("something bad happened")
	defer fmt.Println("end")  //panics are happened after defer statements execute
}


//Recover is a function provided by Go which takes control of a panicking goroutine. 
//recover() can only be used inside deferred functions. 
//If you call recover() during normal flow, it will simply return nil . 
//However, if a goroutine is panicking, recover() will capture the panic value
 
func deferpanicrecover() {
  defer func(){
	  if err := recover(); err != nil {
		  log.Println("Error:",err)
	  }
  }()
  panic("Something bad happend here")
  fmt.Println("end")
  // in this function panic will stop the execution but it can't stops the main function to execute
}