# Prerequisites

- **Go 1.18 or later**

# Introduction

## Non-Generics

In non-generics, if we want our function takes two type as parameter, so basically we create two function, first is we create function it take `int64` and second is `float64`. examples:

```go
// SumInts adds together the values of m.
// It takes type of value is int64.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
// It takes type of value is float64.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
```

Function `SumInts` take paramater `map[string]int64`, so their `key` should be a `string` and the `value` should be `int64`. And function `SumFloats` take paramater `map[string]float64`, so their `key` should be a `string` and the `value` should be `float64`. Now, the problem is, if we want `Sum` another data type example `int32` or `int8` so we create new function that take `int32` or `int8` as `value` in map. So it will be like this:

```go
// This take int32 as value in map.
func SumInt32(m map[string]int32) int32{
    ...
}

// This take int8 as value in map.
func SumInt8(m map[string]int8) int8{
    ...
}
```

## Generics

Thankfully in Go 1.18 or later, we handle those situation with `generics`. We can take any parameter in a function. We can use `comparable`, `any` or you can custom your own. So those code become more simple.

<br>

### constraint `comparable`

---

```go
// Sum can take two data types int64 or float64
func Sum[K comparable, V int64 | float64](m map[K]V) V {
    ...
}
```

Cool right? One function can handle two data types.
`comparable` is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). And map `value` only takes two data types, `int64` and `float64`.

<br>

### constraint `any`

---

With `any` we can't handle sum, cause `any` is an alias for `interface{}` which allows any type. If a parameter can be of any type, that basically won't allow you to do anything with it because you have no guarantee what it will be. Here example code with `any`:

```go
func get[K comparable, V any](m map[K]V, key K) V {
    return m[key]
}
```

The `K` key type must be `comparable`, else it cannot be used as the key type of some map (`m[K]V` in the example). `V` on the other hand shouldn't be constrained, the value type may be anything, and we're not doing anything with it (just returning a value of type `V`), so using `any` here is the best choice.

<br>

### custom constraint

---

We can custom our constraint, and I think this way make code more beauty than the first one.

```go
type Numbers interface {
    // anything data type you want
    int64 | float64 | int32 | int8
}

func Sum[K comparable, V Numbers](m map[K]V) V {
    ...
}
```

<br>

---

The `comparable` constraints only allows types that are comparable, that is, the `==` and `!=` operators are allowed to use on values of them. This is good if you want to use the type as a key in a map (maps require key types to be comparable), or if you you want to find an element in a slice, and you want to use the `==` operator to compare the elements to something.

The `K` key type must be `comparable`, else it cannot be used as the key type of some map (`m[K]V` in the example). `V` on the other hand shouldn't be constrained, the value type may be anything, and we're not doing anything with it (just returning a value of type `V`), so using `any` here is the best choice.

Another example, a slice-find function:

```go
func find[V comparable](what V, s []V) int {
    for i, v := range s {
        if v == what {
            return i
        }
    }
    return -1
}
```

`find()` returns the index of the first occurrence of `what` in the slice `s`, and if it's not in the slice, returns `-1`. The type parameter `V` here must be `comparable`, else you couldn't write `v == what`, using `V any` would be a compile-time error. The constraint `comparable` ensures this `find()` function can only be instantiated with types (and called with values) where the `==` operator is defined and allowed.
