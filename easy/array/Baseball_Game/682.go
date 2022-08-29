package Baseball_Game

import "strconv"

// An integer x - Record a new score of x.
// "+" - Record a new score that is the sum of the previous two scores. It is guaranteed there will always be two previous scores.
// "D" - Record a new score that is double the previous score. It is guaranteed there will always be a previous score.
// "C" - Invalidate the previous score, removing it from the record. It is guaranteed there will always be a previous score.
// Input: ops = ["5","2","C","D","+"]
// Output: 30
// Explanation:
// "5" - Add 5 to the record, record is now [5].
// "2" - Add 2 to the record, record is now [5, 2].
// "C" - Invalidate and remove the previous score, record is now [5].
// "D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
// "+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
// The total sum is 5 + 10 + 15 = 30.
func calPoints(ops []string) int {
	newScore := []int{}
	total := 0
	for _, v := range ops {
		switch v {
		case "+":
			// added previous tow scores
			leng := len(newScore)
			new := newScore[leng-1] + newScore[leng-2]
			newScore = append(newScore, new)
		case "C":
			//delete last element
			newScore = newScore[:len(newScore)-1]
		case "D":
			// add new element to slice from last element double value
			double := newScore[len(newScore)-1] * 2
			newScore = append(newScore, double)
		default:
			score, _ := strconv.Atoi(v)
			newScore = append(newScore, score)
		}
	}
	for _, t := range newScore {
		total += t
	}
	return total
}
