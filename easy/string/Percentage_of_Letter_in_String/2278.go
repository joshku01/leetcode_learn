package Percentage_of_Letter_in_String

// Input: s = "foobar", letter = "o"
// Output: 33
func percentageLetter(s string, letter byte) int {
	count := 0
	totalLen := len(s)
	// range 迴圈預設是由rune類型來編碼,rune類型用來表示utf8字符
	for _, v := range s {
		if v == rune(letter) {
			count++
		}
	}
	percent := float64(count) / float64(totalLen)

	return int(percent * 100)
}
