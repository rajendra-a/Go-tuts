package main

import (
	"fmt"
)

func main(){
	n := 1 == 1 // double equals operator called equals operator it checks 
	m := 1 == 2 //item on the left is equal to right side item or not
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, n)
	var p bool  // every time we initilize variable it has value 0 so here p have 0
	print(p)   // so o means False

	b := 3 //011//we have wide range integer types in the int data types 
	// those are like signed and unsigned suppose uint and int 
	// uint8, uint16, uint32 we don't have 64 bit we have byte which alias for uint8
	// int8, int16, int32, int64 
	fmt.Printf("%v, %T\n", b, b)
	var i uint8 = 45
	fmt.Printf("%v, %T\n", i, i)

	a := 10 //1010
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // we can not get flaoting point when we are using int types
	fmt.Println(a % b) 

	// if we want to add two integers those are not in the same type we have convert into types
	
	//logical operators
	fmt.Println(a & b) // we will get 0010 in binary that is 2
	fmt.Println(a | b)  // we will get 1011 in binary that is 11
	fmt.Println(a ^ b)	// 1001 means 9
	fmt.Println(a &^ b) // 0111 means 8
	
	fmt.Println(a << 3) //64
	fmt.Println(a >> 3) //1

	// there only float 32 and float64 and in complex64 and complex128

	var y complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", real(y), real(y)) 
	fmt.Printf("%v, %T\n", imag(y), imag(y))

	// if we use complex64 it gives float32 if we use complex128 it gives float64
	s := "this is string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2]) 
	fmt.Printf("%v, %T\n", string(s[2]), s[2])// strings are just aliases for bytes so we can get that element
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s + s2, s+ s2)
	c := []byte(s) //converting collection of bytes
	fmt.Printf("%v, %T\n", c, c) // we will get ascii or utf values for each string
	// In go we a data type called rune while string represent any utf-8 character but rune represents utf 32 char
	r := 'a' // note string only represent with double quotes while runes wiht single quotes
	fmt.Printf("%v, %T\n", r, r) //97 we will get data type as int32 so runes as aliases for int32

	var g rune = 'a'
	fmt.Printf("%v, %T", g, g) 
} 