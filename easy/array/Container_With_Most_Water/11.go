package Container_With_Most_Water

import (
	"fmt"
	"sort"
)

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// 題義:找出兩點最高間的數字,畫成圖找出次高的點間可裝水的容量,上題8---->75之間的容量
func maxArea(height []int) int {
	comp := []int{}
	for _, v := range height {
		comp = append(comp, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(comp)))
	max := 0
	second := 0
	for i := 0; i < len(comp); i++ {
		if height[i] > max {
			max = height[i]
			fmt.Println("---->here")
		}
		if height[i] < max {
			second = height[i]
			break
		}
	}
	fmt.Println("------>max", comp)
	fmt.Println("---->second", second)
	return 0

}
