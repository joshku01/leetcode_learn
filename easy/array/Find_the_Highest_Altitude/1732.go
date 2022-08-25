package Find_the_Highest_Altitude

// Input: gain = [-5,1,5,0,-7]
// Output: 1
// Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
func largestAltitude(gain []int) int {
	newArr := make([]int, len(gain)+1)
	newArr[0] = 0
	newArr[1] = gain[0]
	bigist := newArr[0]
	if newArr[1] > newArr[0] {
		bigist = newArr[1]
	}
	for i := 2; i <= len(gain); i++ {
		newArr[i] = newArr[i-1] + gain[i-1]
		if newArr[i] > bigist {
			bigist = newArr[i]
		}
	}
	return bigist
}
