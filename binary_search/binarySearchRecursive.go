package main

import "fmt"

func binarySearchRecursive(arr []int, x int, left int, right int) bool {
	if left > right {
		return false
	}

	mid := (right + left) / 2
	if x == arr[mid] {
		return true
	} else if x < arr[mid] {
		return binarySearchRecursive(arr, x, left, mid-1)
	} else {
		return binarySearchRecursive(arr, x, mid+1, right)
	}
}

func main() {
	a := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	result := binarySearchRecursive(a[:], 30, 0, len(a)-1)
	fmt.Println(result)
}
