package exer9

import (
	"math/rand"
)

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort
	for i, elem := range arr {
		sortedIndex := i-1
		target := elem
		targetIndex := i
		for j := sortedIndex; j >= 0; j-- {
			if target < arr[j] {
				arr[targetIndex], arr[j] = arr[j], arr[targetIndex]
				targetIndex = j
			}
		}
	}
}

const insertionSortCutoff = 0

func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on a good value for insertionSortCutoff
	arrLen := len(arr)
	if arrLen <= insertionSortCutoff {
		InsertionSort(arr)
	} else {
		if arrLen > 1 {
			pivotIndex := arrLen/2
			leftIndex, rightIndex := 0, arrLen-1
			arr[pivotIndex], arr[rightIndex] = arr[rightIndex], arr[pivotIndex]
			for i := range arr {
				if arr[i] < arr[rightIndex] {
					arr[leftIndex], arr[i] = arr[i], arr[leftIndex]
					leftIndex = leftIndex + 1
				}
			}
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
			QuickSort(arr[:leftIndex])
			QuickSort(arr[leftIndex+1:])
		}
	}
}

