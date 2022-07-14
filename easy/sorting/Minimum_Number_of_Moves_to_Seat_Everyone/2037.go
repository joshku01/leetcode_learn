package Minimum_Number_of_Moves_to_Seat_Everyone

import (
	"math"
	"sort"
)

//Input: seats = [3,1,5], students = [2,7,4]
//Output: 4
func minMovesToSeat(seats []int, students []int) int {

	sort.Ints(seats)
	sort.Ints(students)

	seatLen := len(seats)
	var total int
	for i := 0; i < seatLen; i++ {
		diff := int(math.Abs(float64(students[i] - seats[i])))
		total = total + diff
	}
	return total
}
