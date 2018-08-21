package algo

/**
** Approach: Divid and Conquer
**		- In a sorted list, start searching from the middle of the list.
** 		- Then constantly reduce the amount of data that need to be searched.
** Prerequisite: A sorted list.
** Complexity: O(log n)
**/

// BinarySearch find the key in a sorted list. Implemented with loop.
func BinarySearch(data []int, key int) int {
	low := 0
	high := len(data) - 1
	for low < high {
		mid := (high - low) / 2
		if key == data[mid] {
			return mid
		}

		if key > data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearch2 is an implementation with recursion
func BinarySearch2(data []int, key, low, high int) int {
	if high < low {
		return -1
	}

	mid := (high - low) / 2
	if key == data[mid] {
		return mid
	}

	if key > data[mid] {
		return BinarySearch2(data, key, low, mid-1)
	}

	return BinarySearch2(data, key, mid+1, high)
}
