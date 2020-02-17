package main
import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Interfaces")
	var w Writer  = ConsoleWriter{} // w holds the object that implements writer interface
	w.Write([]byte("Hello interfaces"))  // even we don't know what is concreate type but we know how to call  
}

//Interfaces don't describe data it'sdescribes behaviour
type Writer interface {
	//instead of entering bunch of data types we are gonna be stored method definitions
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

// we are gonna creating method on our console writer that has the signature of our writer interface

func (cw ConsoleWriter) Write(data []byte) (int, error) { //concreate method
	n, err := fmt.Println(string(data))
	return n, err
}