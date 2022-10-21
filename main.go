package main

import "fmt"

func task(nums []int) int {
	var left, total int

	for _, num := range nums {
		total += num
	}

	for i, num := range nums {
		total -= num

		if left == total {
			return i
		}

		left += num
	}

	return -1
}

func main() {
	f := []int{1, 7, 3, 6, 5, 6}
	s := []int{1, 2, 3}
	t := []int{2, 1, -1}

	fmt.Println(task(f)) // 3
	fmt.Println(task(s)) // -1
	fmt.Println(task(t)) // 0
}
