package algo

/**
** Approach:
**		- Find the smallest element from the original collection by enumerating every elements in the collection.
**		- Add it to a new collection.
**		- Remove the element from the original collection.
**		- Repeat the steps until all elements are removed from the original collection.
** Complexity: O(n^2)
**/

// SelectionSort sort the input by putting the elements to a new collection in order.
func SelectionSort(data []int) []int {
	var sorted []int
	for len(data) > 0 {
		i := findSmallest(data)
		sorted = append(sorted, data[i])
		data = append(data[:i], data[i+1:]...) // Remove the element from the collection.
	}

	return sorted
}

func findSmallest(data []int) int {
	si := 0 // The index of the smallest element
	for i, v := range data {
		if data[si] > v {
			si = i
		}
	}

	return si
}

// SelectionSort2 implement selection sort without new slice.
func SelectionSort2(data []int) []int {
	for i := 0; i < len(data); i++ {
		s := i
		for j := i; j < len(data); j++ {
			if data[s] > data[j] {
				s = j // Find the smallest
			}
		}
		data[i], data[s] = data[s], data[i] // Swap i and smallest
	}

	return data
}
