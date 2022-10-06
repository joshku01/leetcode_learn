package Maximum_Units_on_a_Truck

// Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
// Output: 91
func maximumUnits(boxTypes [][]int, truckSize int) int {
	var total int
	// 先做排序
	for i := 0; i < len(boxTypes)-1; i++ {
		for j := 0; j < len(boxTypes)-i-1; j++ {
			if boxTypes[j][1] < boxTypes[j+1][1] {
				temp := boxTypes[j]
				boxTypes[j] = boxTypes[j+1]
				boxTypes[j+1] = temp
			}
		}
	}
	// 依序拿出箱子對應的數
	for k := range boxTypes {
		if boxTypes[k][0] <= truckSize {
			temp := boxTypes[k][0] * boxTypes[k][1]
			total += temp
			truckSize -= boxTypes[k][0]
		} else {
			temp := truckSize * boxTypes[k][1]
			total += temp
			break
		}
	}
	return total

}
