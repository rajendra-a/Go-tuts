package main
import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 { // in go increamental operator is statement not an expression 
		fmt.Println(i)					// we need to fugure out the way to do both of these at one time
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
    i := 0 
	for ; i < 5; i++ {  // we must have semicolon
		fmt.Println(i)
	}

	l := 0 
	for l < 5 { 
		fmt.Println(l)
		l++
	}

	m := 0 
	for {  
		fmt.Println(m)
		m++
		if m == 5 {
			break
		}
	}

	 
	for k := 0; k < 10; k++ {  // we must have semicolon to separate the statements
		if k%2 == 0 { 
		continue
	}
	fmt.Println(k)
}  

	for s := 0; s <= 3; s++ {
		for w:= 0; w <= 3; w++{
			fmt.Println(s*w)
		}
	} 
	
loop:	for q := 0; q <= 3; q++ {  // labelling
		for p:= 0; p <= 3; p++{
			fmt.Println(q*p)
			if q * p >= 3 {  // we will get out of the inner loop but it won't exit outer loop so infinate will happened
				break loop   // (if we want to exit outer loop) just apply label with break where we want to exit
			}
		}
	} 

z := []int{1, 2, 3}
for k, v := range z {
	fmt.Println(k, v)
	}


// if we want only keys from collection then
hj := []int{1, 2, 3}
for x := range hj {  // we will get only keys 
	fmt.Println(x)
  }

//if we want to get the values from collection then
gh := []int{4, 5, 6}
for _,value := range gh {  // remember _ is write only operator
	fmt.Println(value)
}
}