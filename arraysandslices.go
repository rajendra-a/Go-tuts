package main

import (
		"fmt"
)

	

func main() {

	//arrays
	grades := [...]int{97, 85, 93} //an array can store one type of data and here we assigned values 
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string  //declaring the array with size of 3 starting index 0
	students[0] = "Lisa"  // assigning to 0 index value	
	students[1] = "ahmed"
	students[2] = "arnold"
	fmt.Printf("students: %v\n", students) 
	fmt.Printf("student #1: %v\n", students[1]) // retrieving the second value from array which is at index 1
	fmt.Printf("length of the students: %v\n", len(students))


	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3} // here  ... means just large enough to hold data that gonna passed in literal syntax
	b := a  // copies a values to b, means if we change the b value that doesn't effect the varable
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	c := &a  // c is pointing to a
	fmt.Println(c) // now c pointing to the a data means both have same address


	//slices
	d := []int{7, 8, 9}
	e := d   // in slices it references the assigned value
	e[2] = 6  // changes index 2 value as well as changes in index2 value in d 
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("Length %v\n", len(d))
	fmt.Printf("capacity %v\n", cap(d))

	f := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	g := f[:] //slice of all elements
	h := f[3:] //slice from 4th element to end
	i := f[:6] //slice first 6 elements
	j := f[3:6] //slice the 4th, 5th and 6th element  
	//slice source is array so we can perform these operations on arrays too
	
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	//we can also use make function to create slices
	k := make([]int, 3, 100) // 2 parameter lenght, 3 parameter is capacity //slices aren't fixed
	fmt.Println(k)
	fmt.Printf("Lenght: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))


	l := []int{}
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("capacity: %v\n", cap(l))
	l = append(l, 1)  // slices aren't fixed 
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("capacity: %v\n", cap(l))
	l = append(l, 2, 3, 4, 5, 6)  // append is variadic function so we can append two or more arguments	
	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("capacity: %v\n", cap(l))

	// if we want add a slice to another slice
	l = append(l, []int{7, 8, 9}...) //this is called spread operation here three dots are spread the slice out into individual arguments	fmt.Println(l)
	fmt.Printf("Length: %v\n", len(l))
	fmt.Printf("capacity: %v\n", cap(l))

	//stack operation how to pop operation
	m := l[1:] //removing first element from slice 

	n := l[ :len(l)-1] // removing last element from slice
	fmt.Println(m)
	fmt.Println(n)

	o := append(l[:4], l[5:]...)
	fmt.Println(o) 
	fmt.Println(l)  //because we have reference to l so make sure that don't have an underlying array 

	



}