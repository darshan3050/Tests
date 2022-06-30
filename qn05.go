// Fill in the blanks so this program compiles and produces
// the output shown.
package main

import "fmt"

func main() {
	// Declare a variable that holds a map with strings for keys
	// and booleans for values.
	// var isVowel map
	// Call make to create the map.
	isVowel := make(map[string]bool)
	//Assign true to the key "a", and false to the keys "b" and "c".
	isVowel["a"] = true
	isVowel["b"] = false
	isVowel["c"] = false
	fmt.Printf("%#v\n", isVowel) // => map[string]bool{"a":true, "b":false, "c":false}

	// Create a map without make with ints as keys and strings as values. The 1 key
	// should have the value "H", 2 should have the value "He", and
	// 3 should have the value "Li".

	elements := make(map[int]string)

	elements[1] = "H"
	elements[2] = "He"
	elements[3] = "Li"
	// Print all the keys and corresponding values in the slice.
	// Order doesn't matter.
	for atomicNumber, value := range elements {
		fmt.Println(atomicNumber, value)
	}
	// => 3 Li
	// => 1 H
	// => 2 He
}