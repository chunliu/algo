package algo

import (
	"bytes"
	"strconv"
)

/**
** Problem: https://www.lintcode.com/problem/add-strings/description
** Approach: Use stack
** Complexity: O(n)
**/

// AddString add the numbers of two string without direct conversion
func AddString(num1, num2 string) string {
	s1, s2, s3 := NewStack(), NewStack(), NewStack()
	c1, c2 := []rune(num1), []rune(num2)

	for _, c := range c1 {
		n, _ := strconv.Atoi(string(c))
		s1.Push(n)
	}

	for _, c := range c2 {
		n, _ := strconv.Atoi(string(c))
		s2.Push(n)
	}

	carry := 0
	for s1.Len() > 0 || s2.Len() > 0 {
		n1, n2 := 0, 0
		if s1.Len() > 0 {
			n1 = s1.Pop().(int)
		}
		if s2.Len() > 0 {
			n2 = s2.Pop().(int)
		}
		n3 := n1 + n2 + carry
		carry = n3 / 10
		s3.Push(n3 % 10)
	}
	if carry > 0 {
		s3.Push(carry)
	}

	var b bytes.Buffer
	for s3.Len() > 0 {
		n := strconv.Itoa(s3.Pop().(int))
		b.WriteString(n)
	}

	return b.String()
}
