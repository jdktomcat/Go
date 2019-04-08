package main

import (
	"fmt"
)

func bubbleSort(array []int) {

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i+1] < array[i] {
				Swap(array, i, i+1)
				swapped = true
			}
		}
	}
}

func Swap(array []int, i, j int) {
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

func main() {

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", array)
	bubbleSort(array)
	fmt.Println("Sorted array: ", array)
}
