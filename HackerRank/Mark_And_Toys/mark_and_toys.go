package Mark_And_Toys

func BubbleSort(array []int32) []int32 {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// input price [1 12 5 111 200 1000 10], k=50
func maximumToys(prices []int32, k int32) int32 {
	// Write your code here
	sorArr := BubbleSort(prices)
	//sort.Slice(prices, func(i, j int) bool {
	//	return prices[i] < prices[j]
	//})
	var total int32
	var mark int32
	for _, v := range sorArr {
		total = total + v
		if total <= k {
			mark++
		}
	}
	return mark
}
