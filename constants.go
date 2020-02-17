package main

import (
		"fmt"
		//"math"
)
const myConst int16 = 44

//enumerated constants
	// most of the time we use in package level but we could use in function level if that made sense application
	const h = iota // iota is counter that we can use when we are creating enumerating constants

	const (
			i = iota
			j = iota
			k = iota
					) //

	const (
			m = iota
			n
			o
					)

	const (
			catSpecialist = iota
			dogSpecialist 
			rabbitSpecialist
	)

	const(
			isAdmin = 1 << iota
			isHeadquaters  //10
			canSeeFinancails //11

			canSeeAfrica //100
			canSeeAsia //101
			canSeeEurope //110
			canSeeNorthAmerica //111
			canSeeSouthAmerica	// 1000 
	)
func main() {
	//const myConst int = 42
	//myConst = 27 // compile throws an error con't assign to myConst becuase we can allowed to change the value of constant 
	fmt.Printf("%v, %T\n", myConst, myConst)
	/*const sinConst float64 = math.Sin(1.57)  require function to execute but that is not allowable at compile time
	fmt.Printf("%v, %T\n", sinConst, sinConst)  so we can't set const that has to be determined at runtime*/

	//constant can be any primitive types

	const x 
	x = float32(43)
	const a int = 14
	const b string = "foo"
	const c float32 = 1.54
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	// constants can be shadowed as variables inner constant shadowed outer constant
	// this is not good shadowing constatn however if situtation comes we can use it

	// now try add constant and var 
	var e int = 24
	fmt.Printf("%v, %T\n", a + e, a + e) //  we can get since they are same type

	//we can do implicit conversions when we working with constants
	const f = 45
	fmt.Printf("%v, %T\n", f, f)
	 //we can add a variable to it
	var g int16 = 45
	fmt.Printf("%v, %T\n", f+g, f+g) 
	/* value replacing every instace with the value and literal 45 adding to an extint to it 
	it interprets as int16 */
	// this can't be done with variables means implicit conversions //


	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	// when we get tha result iota is changing its value as the constants are being evaluated
	// we can take another advantage of iota  if we dont assign the value of constant after the first one
	//then the compiler gonna trying to infer the pattern of the assignments 

	fmt.Printf("%v\n", m) // we would expect have an error  becuase n o haven't value assigned
	fmt.Printf("%v\n", n) //but since we were established a pattern for how to name the constants in the block 
	fmt.Printf("%v\n", o) //when we run we actually get the same result and thats compiler gonna applies the formula 
	// iota is scoped to the block what thats lets us to do we can actullay created related constants together ensure that they have
	// different values and then if you have an other related set of constants we can use another block ensure that they have unique values
	// but allows duplicaion between values in one const block and other
	
	var specialistType int = dogSpecialist  // 
	fmt.Printf("%v\n", specialistType == dogSpecialist )
	//it make sense

	// but when we declare a variable but dont initialize a type
	var specialistType1 int
	fmt.Printf("%v\n", specialistType1 == dogSpecialist)
	// we get false which make sense
	// remember the initial value of iota is 0
	fmt.Printf("%v\n", specialistType1 == catSpecialist)
	//we get true because iota initial value 0 for catspecialist and varialbe with declaration takes stores value 0 that makes no sence 
	//we can use couple of methods for that loop hole
	// we can use first constant value as an error value if our application feels it garbage but we can 
	// use that as an write only variable that is _ symbol
	// it generates value but compiler don't care what it is and It will throw away. remember when we are working with constants value of constants able to remember at compile time
	// but some operations go is gonna allows us to do we can do addition with iota like iota+5	

	// we can also use bit shifting inorder to set boolean flags inside of single byte
	var roles byte = isAdmin | canSeeFinancails | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin & roles == isAdmin)  //bit mask
	fmt.Printf("Is HQ? %v\n", isHeadquaters & roles == isHeadquaters) 
}