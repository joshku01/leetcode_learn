package Decode_the_Message

import (
	"fmt"
	"strings"
)

// Input: key = "the quick brown fox jumps over the lazy dog", message = "vkbs bs t suepuv"
// Output: "this is a secret"
func decodeMessage(key string, message string) string {
	di := "abcdefghijklmnopqrstuvwxyz"
	splDI := strings.Split(di, "")
	mapDi := make(map[string]bool)
	mapRes := make(map[string]string)
	out := ""
	sp := strings.Split(key, " ")
	for _, v := range sp {
		out = out + v
	}
	spStr := strings.Split(out, "")

	var result []string
	for _, str := range spStr {
		if _, ok := mapDi[str]; !ok {
			mapDi[str] = true
			result = append(result, str)
		}
	}
	fmt.Println("---->result", result)
	for k, v := range result {
		fmt.Println("--->v", result[k])
		mapRes[v] = splDI[k]
	}
	fmt.Println("----->map", mapRes)
	splitMessage := strings.Split(message, " ")
	var output []string
	for _, v := range splitMessage {
		var temp string
		fmt.Println("---->", v)
		spliSection := strings.Split(v, "")
		for _, v2 := range spliSection {
			temp += mapRes[v2]
		}
		output = append(output, temp)
	}
	fmt.Println("----->put", output)

	return out

}
