package main

import "fmt"

func main() {

	if true {
		fmt.Println("hi this case passed boolean true")
	}
	statePopulations := map[string]int {
		"California": 39034554,
		"Texas": 43445443,
		"Florida": 54954545,
		"New York": 45435053,
	}

	if pop, ok :=statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	number := 54
	guess := 54
	if guess < 1 || guess > 100 {
		fmt.Println("the guess must between 1 and 100")
	} else if guess < number {
		fmt.Println("Too low")
	}

	if guess > number {
		fmt.Println("Too high")
	}

	if guess == number {
		fmt.Println("you got it!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)


	//Switch
	switch 2 {      //
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("Two or three")
	default:
		fmt.Println("not one r two")
	}

	//Switch
	switch i := 6 + 3 ;i {      //
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("Two or three")
	default:
		fmt.Println("not one r two")
	}


	// Swith with tag less syntax
	i := 11
	switch {      //
	case i <= 10:
		fmt.Println("less than ten")
	case i <= 20:
		fmt.Println("less than twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Swith with tag less syntax
	j := 11
	switch {      //
	case j <= 10:
		fmt.Println("less than ten")
		fallthrough  // it doesn't matter the following case execute because of fallthrough keyword
	case j <= 20:
		fmt.Println("less than twenty")
	default:
		fmt.Println("greater than twenty")
	}
	
	var k interface{} = [3]int{}
	switch k.(type) {
	case int:
		fmt.Println("i is an int ")
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is  string")
	case [3]int:
		fmt.Println("i is [3]int") // in some situations we don't want to execute 
		break
		fmt.Println("this will print too")
	default:
		fmt.Println("i is integer")
	}


}