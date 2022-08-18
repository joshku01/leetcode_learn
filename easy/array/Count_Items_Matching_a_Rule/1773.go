package Count_Items_Matching_a_Rule

// rule [type,color,name]
// Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
// Output: 1
func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	var count int
	rule := map[string]int{"type": 0, "color": 1, "name": 2}
	for _, item := range items {
		index := rule[ruleKey]
		if item[index] == ruleValue {
			count++
		}
	}
	return count
}
