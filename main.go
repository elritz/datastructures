package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	index := binarySearch(a, 11)
	fmt.Println("index =>", index)

}

func binarySearch(a []int, value int) int {
	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (high + low) / 2

		if value == a[mid] {
			return mid
		} else if value < a[mid] {
			high = mid - 1
		} else if value > a[mid] {
			low = mid + 1
		}

	}

	return -1
}
