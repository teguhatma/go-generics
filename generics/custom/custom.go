package custom

import "fmt"

// Custom constraint
type Numbers interface {
	int64 | float64	
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Numbers](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Print(){
	// Initialize a map for the integer values
	ints := map[string]int64{"first": 34, "second": 12}

	// Initialize a map for the float values
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n", 
	SumNumbers(ints),
	SumNumbers(floats))
}