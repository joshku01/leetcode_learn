package Check_if_the_Sentence_Is_Pangram

import (
	"strings"
)

// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// 題意: 給定參數字串,是否全部包含英文字母26個
func checkIfPangram(sentence string) bool {
	en := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	split := strings.Split(sentence, "")
	count := 0
	for _, v := range split {
		for k, e := range en {
			if v == e {
				// 取要刪掉的前面array + 欲刪掉的後面array
				en = append(en[:k], en[k+1:]...)
				count++
			}
		}
	}
	if len(en) == 0 {
		return true
	}
	return false
}
