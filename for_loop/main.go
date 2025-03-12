package main

import "fmt"

func main() {
	// Slice
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

	// Map
	maps := map[string]int{"a" : 1, "b" : 2, "c" : 3}
	for key, val := range(maps) {
		fmt.Println(key, val)
	}

	// String 
	str := "helloworld"
	for idx, char := range str {
		fmt.Println(idx, string(char))
	}
}