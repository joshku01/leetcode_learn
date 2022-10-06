package Can_Place_Flowers

// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true
func canPlaceFlowers(flowerbed []int, n int) bool {

	// if flowerbed is [0,0,0,0,1,0,0,1,0,0,0,0,1,1,0,0,1]

	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if (i > 0 && flowerbed[i-1] == 1) || (i+1 < len(flowerbed) && flowerbed[i+1] == 1) {
			continue
		}
		flowerbed[i] = 1
		n--
		if n == 0 {
			return true
		}
	}
	return false

}
