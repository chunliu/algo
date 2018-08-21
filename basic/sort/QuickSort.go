package algo

import "math/rand"

/**
** Approach: Divid and Conquer
**		- Pick a pivot
**		- Partition the slice into 2 sub-slices: elements less than the pivot, and elements greater than the pivot
**		- Call quick sort recursively on the 2 sub-slices.
**		- Assemble the result: less + pivot + greater
** Complexity: average case: O(n*log n); worst case: O(n^2)
**		- The height of the call stack is O(log n) (average case, worst case is O(n))
**		- and each call stack takes O(n) time.
**/

// QuickSort sort the data by dividing them into two sub slices.
func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	pivot := data[0] // Pick up a random element would improve the performance.
	var less, greater []int
	for i := 1; i < len(data); i++ {
		if data[i] <= pivot {
			less = append(less, data[i])
		} else {
			greater = append(greater, data[i])
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, pivot), greater[0:]...)
}

// QuickSort2 doesn't create any new slice.
func QuickSort2(data []int) []int {
	if len(data) < 2 {
		return data
	}

	left, right := 0, len(data)-1

	pivot := rand.Int() % len(data)

	data[pivot], data[right] = data[right], data[pivot]

	for i := range data {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	QuickSort2(data[:left])
	QuickSort2(data[left+1:])

	return data
}
