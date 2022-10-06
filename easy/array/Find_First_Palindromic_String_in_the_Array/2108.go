package Find_First_Palindromic_String_in_the_Array

import "fmt"

// Input: words = ["abc","car","ada","racecar","cool"]
// Output: "ada"
// Explanation: The first string that is palindromic is "ada".
// Note that "racecar" is also palindromic, but it is not the first.
func firstPalindrome(words []string) string {
	// 找出奇數還是偶數
	var markWord []string
	for i := 0; i < len(words); i++ {
		var mark int
		for j := 0; j < len(words[i])/2; j++ {
			head := fmt.Sprintf("%c", words[i][j])
			tail := fmt.Sprintf("%c", words[i][len(words[i])-1-j])
			if head == tail {
				mark++
				continue
			} else {
				break
			}
		}
		// 數量正確append到array
		if mark == len(words[i])/2 {
			markWord = append(markWord, words[i])
		}
	}
	if len(markWord) != 0 {
		return markWord[0]
	} else {
		// 沒有則返回空字串
		return ""
	}
}
