package Pascals_Triangle

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
/*跟Fibonacci Sequence 觀念類似
step1.先抓出前兩筆資料
step2.從第三筆開始算出迴圈內的數字
step3.存入一個初始化的變數n2內繼續做計算
*/
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	// 初始化n2
	n2 := [][]int{{1}, {1, 1}}

	for i := 3; i <= numRows; i++ {
		//前一個結果的變化
		// 0-3  0,1,2
		var r []int
		for j := 0; j <= i-1; j++ {
			//取n2的最後一個arr , len(n2)-1
			switch j {
			case 0:
				r = append(r, n2[len(n2)-1][0])
			case i - 1:
				r = append(r, n2[len(n2)-1][0])
			default:
				r = append(r, n2[len(n2)-1][j]+n2[len(n2)-1][j-1])
			}
		}
		n2 = append(n2, r)
	}
	return n2
}
