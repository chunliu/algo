package algo

/**
** Approach:
**		- For data[i], find the first j where j < 1 and data[j] > data[i]
**		- Move data[i] to data[j], and move each data in data[j:i] +1 index to right
** Complexity: O(n^2)
**/

// InsertionSort sort data by inserting the data to the correct position.
func InsertionSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[j] > data[i] {
				data = append(
					append(data[:j],
						append([]int{data[i]}, data[j:i]...)...), data[i+1:]...)
			}
		}
	}

	return data
}
