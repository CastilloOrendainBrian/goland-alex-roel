package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Includes[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(Includes(strings, "a")) // true
	fmt.Println(Includes(strings, "f")) // false
	fmt.Println(Includes(numbers, 1))   // true
	fmt.Println(Includes(numbers, 4))   // false

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))
}
