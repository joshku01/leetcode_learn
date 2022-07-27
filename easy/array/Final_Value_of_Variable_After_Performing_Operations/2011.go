package Final_Value_of_Variable_After_Performing_Operations

var mapStr = map[string]int{
	"++X": 1,
	"X++": 1,
	"--X": -1,
	"X--": -1,
}

// Input: operations = ["--X","X++","X++"]
// Output: 1
func finalValueAfterOperations(operations []string) int {
	var total = 0
	for _, v := range operations {
		total += mapStr[v]
	}
	return total
}
