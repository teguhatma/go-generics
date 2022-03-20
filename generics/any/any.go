package any

import "fmt"

// iterate is iteration a slice with any data types.
func iterate[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// get is to get value from key.
func get[K comparable, V any](m map[K]V, key K) V {
    return m[key]
}

func Print() {
	// Initialize a map for the integer values
	mapInts := map[string]int64{"first": 34, "second": 12}

	sliceString := []string{"data", "types", "are", "wow"}
	sliceInt := []int{1, 2, 3, 4, 5}
	var key string = "data"

	fmt.Println("Integer")
	iterate(sliceInt)

	fmt.Println("String")
	iterate(sliceString)


	result := get(mapInts, key)
	fmt.Println(result)
}
