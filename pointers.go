package main

import "fmt"

func main() {
	var a int = 42  
	var b *int = &a    // declaring variable b with pointer to int 
	fmt.Println(a, b)  // with b we will get numerical representation of memory address that is holding the location of a
	// means here b isn't holding the value 42 it's holding the memory address of the holding a data
	// we can prove that by printing a with address of operator &

	fmt.Println(&a)

	// the values are exactly the same

	// so we can get the value stored in the b adressing location by dereffering operator that is *

	fmt.Println(*b)

	//now change the value of a
	a = 26

	fmt.Println(a, *b)
	// we will get the same value because they are pointing to the same underlying data

	// we can change the value with derefencing operator
	*b = 45

	fmt.Println(a, *b)

	p := [3]int{1, 2, 3}
	q := &p[0] // pointer to the first element in the array
	r := &p[1]  // pointer to the second element in the array
	fmt.Println("%v %p %p\n", p, q, r) // %v for value and %p for pointer for b and c

	s := [3]int{1, 2, 3}
	t := &p[0] // pointer to the first element in the array
	//u := &p[1] - 4 // we may think this for pointer to the first element in the array
	//this won't allow in go pointers 
	// still if we want to perform we could do by importing package calle "unsafe"
	
	fmt.Println("%v %p %p\n", s, t, t) // %v for value and %p for pointer for b and c

	// we can observe & address of operator allowing us to instantiate a pointer to a certain variable
	// if we look at t its actually a pointer to certain variable we had declare the underlying type first
	// this is not neccessary, we often work with only pointers
	type myStruct struct {
		foo int
	}

	// this we can't do this but this code will help us in next block of code
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)  // it says it holds the address of the object which has the field value 42


	
	// we can also use built in new function but we can't do intialization with new function
	// we can create empty object like below

	var ns *myStruct
	fmt.Println(ns) // when we create a new pointer it holds nil value
	ns = new(myStruct)
	fmt.Println(ns)  // it holds 0 as we know that variable declared will hold a value 0
	// now derence and assign a value to it 
   (*ns).foo = 45
   fmt.Println((*ns).foo)
   // this is little ugly to write 
   // for that compiler help us little bit with that we can do as below
   ns.foo = 89
   fmt.Println(ns.foo)


}