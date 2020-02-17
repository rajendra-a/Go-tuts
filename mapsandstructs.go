package main

import (
	"fmt"
	"reflect"
)

//struct
type Doctor struct {
	number int `required max:100`
	actor string
	companions []string
}

func main() {
	statePopulations := make(map[string]int) // declaring map with make function
	statePopulations = map[string]int{
		"California": 32453253, 
		"Texas": 43434287,
		"New York": 54484343,
	} //literal syntax
	
	fmt.Println(statePopulations)
	m := map[[3]int]string{} //array is valid key for map but slice not
	fmt.Println(m)

	
	fmt.Println(statePopulations["California"]) //accessing the value from map

	//adding value to the map
	statePopulations["Georgia"] = 84585044
	fmt.Println(statePopulations["Georgia"])

	//delete Georgia
	delete(statePopulations,"Georgia")
	fmt.Println(statePopulations)

	//now check once agian for georgia 
	fmt.Println(statePopulations["Georgia"]) // we will get 0

	//when we accessing we can add optional ok var
	pop, ok := statePopulations["Georgia"] // when we were not sure if a key is in the map or not
	fmt.Println(pop, ok)  


	//if we want to if we want to just check for present
	// we can use
	_, ok = statePopulations["Georgia"]  // remember _ is write only opearator

	fmt.Println(ok)

	sp := statePopulations
	delete(sp, "California")
	fmt.Println(sp) // we can see now calofornia removed 
	fmt.Println(statePopulations) // we can see that california removed as effected in so map\
	
	aDoctor := Doctor{    // we can also instantiate with positional syntax means no need to mention field names
		number: 3,
		actor: "John",
		companions: []string{
			"liz shaw",
			"jo grant",
		}, 

	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actor)

	lawyer := struct{name string}{name: "sardar"} // there is no type for it so we can call it as anonymous struct
	fmt.Println(lawyer) //and we can't use it because it doesn't have independent name to refer to

	// unline slices and maps these are value types 
	// so whenever we asigned to new variable it copies and doesn't effect to the assigned variable values
	anotherLawyer := lawyer
	anotherLawyer.name = "Tom"
	fmt.Println(lawyer)
	fmt.Println(anotherLawyer)

	// values remain independent
	// if we want to point same data then we can use & address of operator
	lawyer2 := &lawyer
	fmt.Println(lawyer2)

	//embedding
	//go doesn't support object oriented priciples so we don't have inheritance
	//it uses inheritance model called composition

	type Animal struct {
			Name string   `required max:100`
			Origin  string
	}

	type Bird struct {
		Animal   // this is not inheritance its composition means has characteristics of Animal
		Speed	float32   // 
		CanFly	bool 

	}

	b := Bird{} // we declared the type 
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b.Name)

	// if we want to use literal syntax 
	c := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed: 30,
		CanFly: false, 
	}


	fmt.Println(c)

	f := reflect.TypeOf(Doctor{})
	field, _ := f.FieldByName("number")
	fmt.Println(field.Tag)

}