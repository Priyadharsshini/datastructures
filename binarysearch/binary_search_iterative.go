package main

import "fmt"

func binarySearchIterative(arr []int, x int, left int, right int) bool {
	for left <= right {
		mid := (right + left) / 2
		if x == arr[mid] {
			return true
		} else if x < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	a := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	result := binarySearchIterative(a[:], 30, 0, len(a)-1)
	fmt.Println(result)
}
