package algo

/**
** Approach: Divid and Conquer
**		- Equally split the list to 2 sub lists.
**		- Recursively run merge sort on the 2 sub lists.
**		- Merge the sorted sub lists into 1 list.
** Complexity: O(n*log n)
**		- The height of the call stack is O(log n).
**		- and each call stack takes O(n) time to split the list into 2.
**		- Usually quick sort is faster than merge sort.
**/

// MergeSort sort the elements by dividing them to sublists and merging them back to 1 list
// when the sublists are sorted.
func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	mid := len(data) / 2
	var left, right []int
	for i := 0; i < mid; i++ {
		left = append(left, data[i])
	}

	for i := mid; i < len(data); i++ {
		right = append(right, data[i])
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	leftIndex, rightIndex, resultIndex := 0, 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}

	for i := leftIndex; i < len(left); i++ {
		result[resultIndex] = left[i]
		resultIndex++
	}

	for i := rightIndex; i < len(right); i++ {
		result[resultIndex] = right[i]
		resultIndex++
	}

	return
}
