package Flipping_an_Image

// Input: image = [[1,1,0],[1,0,1],[0,0,0]]
// Output: [[1,0,0],[0,1,0],[1,1,1]]
// first flip
// second invert
func flipAndInvertImage(image [][]int) [][]int {
	for _, v := range image {
		lengH := len(v)
		helf := lengH / 2
		for i := 0; i < helf; i++ {
			v[i], v[lengH-1-i] = v[lengH-1-i], v[i]
		}
		for k2, v2 := range v {
			if v2 == 0 {
				v[k2] = 1
			} else {
				v[k2] = 0
			}
		}
	}

	return image

}
