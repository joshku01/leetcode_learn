package Container_With_Most_Water

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// 思考方向:找到兩個數字相差最小,距離差距最遠的
// 利用for迴圈依序慢慢的找,相差距離*相減最小的值存入res中,為容積最大量
func maxArea(height []int) int {

	res := 0
	start := 0
	end := len(height) - 1

	for start < end {
		cur := min(height[start], height[end]) * (end - start)
		if cur > res {
			res = cur
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
