package main

import (
	"fmt"

	any "github.com/teguhatma/go-generics/generics/any"
	comp "github.com/teguhatma/go-generics/generics/comparable"
	"github.com/teguhatma/go-generics/generics/custom"
)


func get(m map[string]int64, key string) int64 {
	return m[key]
}

func main() {
	ints := map[string]int64{
		"first": 12,
		"second": 32,
	}

	result := get(ints, "first")
	fmt.Println(result)

	fmt.Println("This contraint comparable")
	comp.Print()
	fmt.Println()
	fmt.Println("This contraint any")
	any.Print()
	fmt.Println()
	fmt.Println("This custom contraint")
	custom.Print()
}