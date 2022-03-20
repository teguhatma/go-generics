package comparable

import "fmt"

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V
    for _, v := range m {
        s += v
    }
    return s
}

// find is to get the index of value. 
func find[V comparable](what V, s []V) int {
	for i, v := range s {
		if v == what {
			return i
		}
	}
	return -1
}


func Print() {
	var key string = "data"
	var n int = 5
	dataString := []string{"data", "wow"}
	dataInt := []int{1, 2, 3, 4, 5}

	// Initialize a map for the integer values
	ints := map[string]int64{"first": 34, "second": 12}

	// Initialize a map for the float values
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))


	findResult := find(key, dataString)
	findResultInt := find(n, dataInt)
	fmt.Println(findResult)
	fmt.Println(findResultInt)
}
