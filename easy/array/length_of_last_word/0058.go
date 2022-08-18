package length_of_last_word

// for example give string "Hello World" return the length of last word
func lengthOfLastWord(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}
	res := 0
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}

	return res
}

func lengthOfLastWordSecond(s string) int {
	// example string "Hello World"
	start := len(s) -1
	//	寻找最后一个单词的尾下标
	for start >= 0 && s[start] == ' ' {
		start--
	}
	//	寻找最后一个单词的初始下标
	end := start
	for  start >= 0 && s[start] != ' '{
		start--
	}

	return end - start

}
