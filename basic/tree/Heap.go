package tree

/**
** Approach: Use the feature of heap
**	- Make a heap with all items in the array. According to heap's feature, the root item is the biggest.
**	- Swap the root with the last item i in the heap. Then consider item i as being removed from the heap.
**	- Make the rest of i - 1 items to be a heap again. Either recursively call MakeHeap (slower),
**	  or push down the new root with a loop (faster).
** Complexity: O(n*log n)
**	- MakeHeap: O(n*log n) + Push down root: O(n*log n)
**/

// HeapSort sort the array with heap
func HeapSort(data []int) []int {
	// Make a heap
	data = MakeHeap(data)
	for i := len(data) - 1; i >= 0; i-- {
		// Item i is the smallest. Swap it with the root which is the biggest item.
		data[0], data[i] = data[i], data[0]
		// Make the rest of i - 1 items to be a heap again by pushing down the new root
		index := 0
		for true {
			child1 := 2*index + 1
			child2 := 2*index + 2
			// If child index is bigger than the count, there is no more child.
			if child1 >= i {
				child1 = index
			}
			if child2 >= i {
				child2 = index
			}
			// If item index is bigger than both children, it is already a heap.
			if data[index] >= data[child1] && data[index] >= data[child2] {
				break
			}
			// Otherwise, swap the item index with the bigger child. Then repeat to push it further down.
			if data[child1] > data[child2] {
				data[child1], data[index] = data[index], data[child1]
				index = child1
			} else {
				data[child2], data[index] = data[index], data[child2]
				index = child2
			}
		}
	}

	return data
}

// MakeHeap create a heap with array. Complexity: O(n*log n)
func MakeHeap(data []int) []int {
	// Add each item to the heap one at a time.
	for i := 0; i < len(data); i++ { // O(n)
		index := i
		for i > 0 { // O(log n) since the complete binary tree is O(log n) level tall.
			// For item i, find its parent
			parent := (index - 1) / 2
			// If item i <= parent, it meets the heap condition so break.
			if data[index] <= data[parent] {
				break
			}

			// Otherwise, swap the item i and parent, and repeat the steps for new parent value.
			data[index], data[parent] = data[parent], data[index]
			index = parent
		}
	}

	return data
}
