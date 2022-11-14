package sorting

import "fmt"

func BubbleSort(arr []int) {
	leng := len(arr)
	// 總共需要 n 次排序
	for i := 0; i < leng; i++ {
		// 每一回合要比較的次數 n-1
		for j := 1; j < leng-1; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println("---->output", arr)
}
