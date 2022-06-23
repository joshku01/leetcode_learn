package valid_parentheses

import (
	"strings"
)

var r = []string{"(", ")", "{", "}", "[", "]"}

var r1 = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func inArray(val string, array []string) (exists bool) {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

var tag []string

// input "[](){}"
//        012345
// output true
func isValid(s string) bool {
	str := strings.Split(s, "")
	strlen := len(str)
	// check string is validate
	for _, v := range str {
		is := inArray(v, r)
		if !is {
			return false
		}
	}
	isPA := strlen % 2
	if isPA != 0 {
		return false
	}
	for i := 0; i < strlen; i += 2 {
		for k := range r1 {
			if str[i] == k {
				// if not true then tag
				if str[i+1] != r1[k] {
					tag = append(tag, str[i])
				}
			}
		}
	}
	if len(tag) > 0 {
		return false
	}
	return true

}
