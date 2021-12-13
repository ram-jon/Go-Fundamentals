package main

import "fmt"

func main() {
	shoppingCart := map[string]int{
		"Keyboard": 1000,
		"Mouse":    500,
		"Laptop":   50000,
	}

	fmt.Println(shoppingCart)
	// create map using make function
	abc := make(map[string]int)
	abc = map[string]int{
		"adas": 212,
		"sa":   423,
		"Fs":   545,
	}
	fmt.Println(abc)

	//set and get values of map using keys

	shoppingCart["Keyboard"] = 999
	fmt.Println(shoppingCart["Keyboard"])
	fmt.Println(shoppingCart["keyboard"]) //print 0 since key not present in map

	// check if key present in map or not
	price, ok := shoppingCart["keyb"] //the ok variable returns true or false depend on key is present or not
	fmt.Println(price, ok)
	// we can also skip first value using "_" operator
	_, ok = shoppingCart["Keyboard"]
	fmt.Println(ok)

	// copy map to another variable
	// Note: map is reference data type like slices

	sc := shoppingCart

	fmt.Println(shoppingCart)
	fmt.Println(sc)

	sc["Keyboard"] = 500 //it will update Keyboard value in both shoppingCart and sc

	fmt.Println(shoppingCart)
	fmt.Println(sc)

	// Delete key from map

	delete(shoppingCart, "Keyboard")
	fmt.Println(shoppingCart)

}
