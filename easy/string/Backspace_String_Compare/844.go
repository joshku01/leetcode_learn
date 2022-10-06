package Backspace_String_Compare

// Input: s = "ab#c", t = "ad#c"
// Output: true
// Explanation: Both s and t become "ac".
func backspaceCompare(s string, t string) bool {
	if helper(s) == helper(t) {
		return true
	} else {
		return false
	}
}

func helper(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if s[i] == '#' {
			// 減去前面一個
			stack = stack[:len(stack)-1]
		} else {
			// append
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
