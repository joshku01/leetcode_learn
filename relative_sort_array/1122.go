package relative_sort_array

import "fmt"

func relativeSortArray(A, B []int) []int {
	count := [1001]int{}
	for _, a := range A {
		count[a]++
	}
	res := make([]int, 0, len(A))
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}

	return res
}

func result() {
	A := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	B := []int{2, 1, 4, 3, 9, 6}
	res := relativeSortArray(A, B)
	fmt.Println(res)

}
