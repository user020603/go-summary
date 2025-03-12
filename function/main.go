package main

import "fmt"

func printEven(nums ...int) {
	for _, n := range nums {
		if (n & 1) == 0 {
			fmt.Println(n)
		}
	}
}

func cntNum(nums ...int) (int, int) {
	cnt := 0
	for _, n := range nums {
		if (n & 1) == 0 {
			cnt += 1
		}
	}
	return cnt, len(nums) - cnt
}

func main() {
	printEven(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cntNum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}