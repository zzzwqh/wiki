package main

import "fmt"

func main() {

	// Declaring some variables
	var name string
	var alphabet_count int
	var float_value float32
	var boolean_value bool

	// Calling the Sscanf() function which
	// returns the number of elements
	// successfully parsed and error if
	// it persists
	n, err := fmt.Sscanf("GeeksforGeeks 13 6.7 true",
		"%s %d %g %t", &name, &alphabet_count,
		&float_value, &boolean_value)

	// Below statements get executed
	// if there is any error
	if err != nil {
		panic(err)
	}

	// Printing the number of elements
	// and each elements also
	fmt.Printf("%d:%s, %d, %g, %t", n, name,
		alphabet_count, float_value, boolean_value)

}
