package First_Letter_to_Appear_Twice

// Input: s = "abccbaacz"
// Output: "c"
func repeatedCharacter(s string) byte {
	// byte類型是uint8的別名類型,一個值是一個ASCII碼值
	// 字串可以拆解成最小byte單位 ,格式化顯示可以用%q輸出
	// fmt.Printf("%q", []rune(s))
	es := []string{"test"}
	for k, vv := range s {
		for _, e := range es {
			if string(vv) == e {
				return s[k]
			}
			es = append(es, string(vv))
		}
	}

	return s[0]
}
