package CountingBits

import (
	"fmt"
	"strings"
)

// Input: n = 2
// Output: [0,1,1]
func countBits(n int) []int {

	var out []int
	for i := 0; i <= n; i++ {
		b := fmt.Sprintf("%b", i)
		split := strings.Split(b, "")
		count := 0
		for _, v := range split {
			if v == "1" {
				count++
			}
		}
		out = append(out, count)
	}
	return out
}
