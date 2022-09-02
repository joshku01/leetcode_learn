package valid_parentheses

var r1 = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

// input "[](){}"
//        012345
// output true
// step 1 總共有三種符號,每一種至少都有頭跟尾
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[', '{', '(':
			stack = append(stack, s[i])
		case ']', '}', ')':
			if len(stack) == 0 {
				return false
			}
			ch := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 兩邊需要對稱,如果只有一邊有的話return false
			if (s[i] == ']' && ch != '[') || (s[i] == '}' && ch != '{') || (s[i] == ')' && ch != '(') {
				return false
			}
		}
	}
	return len(stack) == 0

}
