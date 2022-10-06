package Best_Time_to_Buy_and_Sell_Stock

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
func maxProfit(prices []int) int {

	// 此方法效能太差 runtime limit
	//diff := 0
	//for i := 0; i < len(prices); i++ {
	//	for j := i + 1; j < len(prices); j++ {
	//		//
	//		if prices[i] < prices[j] {
	//			temp := prices[j] - prices[i]
	//			if temp >= diff {
	//				diff = temp
	//			}
	//		}
	//	}
	//}

	min := 999999
	max := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max

}
