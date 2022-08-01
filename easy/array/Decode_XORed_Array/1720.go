package Decode_XORed_Array

// Input: encoded = [6,2,7,3], first = 4
// Output: [4,2,0,7,4]
func decode(encoded []int, first int) []int {

	// 定義長度=n+1
	res := make([]int, len(encoded)+1)
	// first element= first
	res[0] = first

	// skip res[0] element
	for i, v := range encoded {
		res[i+1] = res[i] ^ v
	}

	return res

}
