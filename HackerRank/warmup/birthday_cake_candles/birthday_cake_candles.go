package birthday_cake_candles

// [4,4,1,3]
// 題目:找出最大高度的蠟燭,然後計算出數量返回
func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here

	max := int64(int64(candles[0]))

	sum := int32(0)
	for i := 0; i < len(candles); i++ {
		if int64(candles[i]) > max {
			max = int64(candles[i])
		}
	}
	// count the mount
	for i := 0; i < len(candles); i++ {
		if int64(candles[i]) == max {
			sum++
		}
	}
	return sum
}
