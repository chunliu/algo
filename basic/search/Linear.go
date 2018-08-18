package algo

/*
Method: enumerate every element in the collection until the key is found.
Complexity: O(n)
*/

// LinearSearch search the key by enumerating every element in the collection.
func LinearSearch(data []int, key int) int {
	r := -1
	for i, v := range data {
		if v == key {
			r = i
			break
		}
	}

	return r
}
