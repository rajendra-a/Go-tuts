package main

import (
	"fmt" 
	"strconv"
)
/* when you want to declare variable you need to write with syntax of specifying
data type */

var l int = 76

var (
	actorname string = "sarah"
	companion string = "vandella"
	dnubmer int = 3
	season int = 11
	
)

var(
	counter int = 0
)
func main() {
	var i int         //variable declaration
	i = 11     
	       //variable initialization
	var j float32 = 9

	var m float32 = 19.6 //variable assignment with data type

	k := 45 //assigning with type inference

	k = 89  // we can reassign 

	//k := 98   we can't redeclare the variable

	fmt.Printf("%v", l) // prints 76 because it takes the variable at package which is already available
	var l int = 99 // redeclared in the function scope which is already in the package scope
	//m := 45  gives error as m is declared and not used which is compile time error
	fmt.Println(i, j, k,"\n")

	fmt.Printf("%v,%T\n", i, i)

	fmt.Printf("%v,%T\n", j, j) //string formatting

	fmt.Printf("%v,%T\n", l, l)

	fmt.Printf("%v\n", l) // prints 99 because it takes value in the function level its called shadowing

	fmt.Printf("%v\n", actorname)

	fmt.Printf("%v, %T\n", m, m)

	var n int 
	n = int(m)
	fmt.Printf("%v, %T\n", n, n)

	var p string
	p = string(n)
	fmt.Print(p,"\n") //  we can not the result as we think we will get the * symbol
	//becuase in go it looks for the ascii character and it gives the that
	//for getting the string type we need to use strconv package for strings
	var q string
	q = strconv.Itoa(n)
	fmt.Printf("%v", q)

}
