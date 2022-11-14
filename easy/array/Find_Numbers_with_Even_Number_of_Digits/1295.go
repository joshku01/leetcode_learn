package Find_Numbers_with_Even_Number_of_Digits

// Input: nums = [12,345,2,6,7896]
// Output: 2
func findNumbers(nums []int) int {

	num := 0

	for i := 0; i < len(nums); i++ {
		target := digitCount(nums[i])
		if target%2 == 0 {
			num++
		}
	}
	return num

}

func digitCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}
