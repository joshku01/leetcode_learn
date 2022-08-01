package Mark_And_Toys

import (
	"fmt"
	"strings"
)

func TimeTrans() map[string]string {
	m := map[string]string{
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
		"12": "12",
	}
	return m
}

func BubbleSort(array []int32) []int32 {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func timeConversion(s string) string {
	// Write your code here
	split := strings.Split(s, ":")
	spLen := len(split)
	getZone := split[spLen-1]
	getZoneType := getZone[len(getZone)-2:]
	trimZoneString := getZone[:len(getZone)-2]
	data := TimeTrans()
	var out string
	if getZoneType == "PM" {
		out = fmt.Sprintf("%s:%s:%s", data[split[0]], split[1], trimZoneString)
	} else {
		if split[0] == "12" {
			out = fmt.Sprintf("%s:%s:%s", "00", split[1], trimZoneString)
		} else {
			out = fmt.Sprintf("%s:%s:%s", split[0], split[1], trimZoneString)
		}
	}

	return out
}

// input price [1 12 5 111 200 1000 10], k=50
func maximumToys(prices []int32, k int32) int32 {
	// Write your code here
	sorArr := BubbleSort(prices)
	//sort.Slice(prices, func(i, j int) bool {
	//	return prices[i] < prices[j]
	//})
	var total int32
	var mark int32
	for _, v := range sorArr {
		total = total + v
		if total <= k {
			mark++
		}
	}
	return mark
}

// input  q = [2, 5, 1, 3, 4]
// output print 3
func minimumBribes(q []int32) []int32 {
	// Write your code here
	totalLen := len(q)
	for k := 0; k < totalLen; k++ {
		for j := k + 1; j < totalLen; j++ {
			temp := 0
			fmt.Println("---->temp", temp)
			if temp > 2 {
				fmt.Printf("Too chaotic")
				break
			}
			if q[k] > q[j] {
				temp++
				q[k], q[j] = q[j], q[k]
			}
		}
	}

	return q
}
