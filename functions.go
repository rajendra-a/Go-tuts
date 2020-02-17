package main

import (
		"fmt"
) 

func main(){
	sayHello("Hello Go")  // calling function with passing argument to parameter

	for i := 0; i < 5 ; i++ {
		sayMessage("Hello Go", i)
	}

	greet("Hello", "Rajendra")


	greeting1 := "hello"
	name := "rajendra"
	Greeting1(greeting1, name)
	fmt.Println(name)

	 s := sum(1, 2, 4, 5, 6)
	 fmt.Println("The result is", s)


	 d, err := devide(5.0, 0.0)
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
	 fmt.Println(d)


	 //anonymous functions

	for i := 0;  i < 5; i++ {  // we aren't reading the outer scope  we are passing into function execution
	func(i int) { 
		fmt.Println(i)
	 }(i)
	}

//we can also work with functions as variables

var f func() = func() {
	fmt.Println("Hello")

}
f()


var devide func(float64, float64) (float64, error) 
devide = func (a, b float64)(float64, error)  {
	if b == 0.0 {
		return 0.0, fmt.Errorf("can't devide by zero")
	} else {

	return a/b, nil
}
d, err := devide(5.0, 0.0)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(d)	
}

g := greeter{
	greeting : "hello",
	name : "rajendra",
}
}

func sayHello(msg string){  // when we can create functions with parameters here msg is parameter and type string
	fmt.Println(msg)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("the value of index is", idx)
}


// we can of when we are defining function we are gonna passing multiple parameters of the same type

func greet(greeting string, name string)  {
	fmt.Println(greeting, name )
	
}

// so when we create parameters with same type we can put together the variables wiht comma delimeter which take care by the compiler

func Greeting1(greeting, name1 string) {
	fmt.Println(greet, name1)
	name1 = "Ted"
	fmt.Println(name1)
}

// variadic parameters

func sum(values ...int) int { 
	// not receiving 5 variables here instead one variable and preceeded its type with three dots 
	// means take all the arguments and wrapped as slice that has name values variadic parameter must be one 
	// that is placed as last parameter 
	fmt.Println(values)
	result := 0
	for _, v := range values{
		result += v
	}
	return result
}

// we can return two or more values
func devide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("can't devide by zero")
	}

	return a/b, nil
}


// methods 
// method just executes the inner known context
type greeter struct {
	greeting string
	name string
}

func (g greeter) greet(){
	fmt.Println(g.greeting, g.name) 
	//greet method gets the object of "greeter" and given to context name "g"
}




