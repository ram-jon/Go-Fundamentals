package main

import "fmt"

func main() {
	// simple for loop to print 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//Using multiple variables inside loop

	for i, j := 1, 5; i <= 5; i, j = i+1, j-1 { //the scope i and j are only inside for block
		fmt.Println(i, j)
	}

	// Using for loop like while
	i, j := 1, 5
	for i <= 5 {
		fmt.Println(i, j)
		i, j = i+1, j-1
	}

	// break and continue inside for
	fmt.Println("---------------------------------------------")

	k, l := 1, 5
	for k <= 5 {
		if l == 1 {
			break
		}
		if k == 2 {
			// continue
		}
		fmt.Println(k, l)
		k, l = k+1, l-1
	}

	// nested for loops
	fmt.Println("*******************************************************")
	for x := 0; x <= 5; x++ {
		for y := 0; y <= 5; y++ {
			if x == 0 || y == 0 {
				continue
			}
			fmt.Println(x * y)
		}
	}

	// iterate over slices, arrays
	slc := []int{1, 2, 3, 4, 5, 6}
	for i, v := range slc { //iterate over array index and value
		fmt.Println(i, v)
	}

	for i := range slc { //iterate over array index  only
		fmt.Println(i)
	}

	for _, v := range slc { //iterate over array value value only
		fmt.Println(v)
	}

	//iterate over maps
	shoppingCart := map[string]int{
		"Keyboard": 1000,
		"Mouse":    500,
		"Laptop":   50000,
	}
	shoppingCart["Monitor"] = 15000

	for k, v := range shoppingCart { //iterate over key value of maps
		fmt.Println(k, v)
	}

}
