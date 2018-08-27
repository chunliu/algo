package ds

/**
** Problem: https://www.lintcode.com/problem/valid-parentheses/description
** Approach: Use stack
** Complexity: O(n)
**/

// IsValidParentheses check if the string contains valid parentheses.
func IsValidParentheses(s string) bool {
	characters := []rune(s)
	st := NewStack()
	for _, c := range characters {
		s1 := string(c)
		if st.Len() == 0 {
			st.Push(s1)
		} else {
			s2 := st.Peek().(string)
			switch string(s2) {
			case "{":
				if s1 != "}" {
					st.Push(s1)
				} else {
					st.Pop()
				}
				break
			case "[":
				if s1 != "]" {
					st.Push(s1)
				} else {
					st.Pop()
				}
				break
			case "(":
				if s1 != ")" {
					st.Push(s1)
				} else {
					st.Pop()
				}
				break
			default:
				return false
			}
		}
	}

	return st.Len() == 0
}
