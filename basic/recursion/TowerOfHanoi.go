package algo

import (
	"fmt"
	"strings"
)

/**
** Problem: https://www.lintcode.com/problem/tower-of-hanoi/description
** Approach: Recusion
**		- Recursively move the n-1 disk from From peg to Other peg.
**		- Move the n-th disk (which is the biggest) from From peg to To peg.
**		- Recursively move the n-1 disk from Other peg to To peg.
** Complexity: O(2^n)
**/

// TowerOfHanoi resolve the Tower Of Hanoi problem.
func TowerOfHanoi(n int) []string {
	var result []string
	result = moveTower("A", "C", "B", n, result)

	return result
}

func moveTower(from, to, other string, n int, result []string) []string {
	if n > 1 {
		// 1. Recursively move the n-1 smaller disk from From to Other
		result = moveTower(from, other, to, n-1, result)
	}

	// 2. Move n-th disk from From peg to To peg.
	var b strings.Builder
	fmt.Fprintf(&b, "from %s to %s", from, to)
	result = append(result, b.String())

	if n > 1 {
		// 3. Recursively move the n-1 disk from Other peg to To peg.
		result = moveTower(other, to, from, n-1, result)
	}

	return result
}
