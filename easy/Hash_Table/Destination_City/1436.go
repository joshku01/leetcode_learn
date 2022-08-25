package Destination_City

// Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
// Output: "Sao Paulo"
// 題意: 在paths的列表內 找出一個位置陣列arr[1]是為終點,即無法連接到下一個城市的開頭
// 解題步驟:
// step1: 先用map方式列出所有的城市,並設定為0
// step2: 在映射出在第二個位置是否有出現過,有出現則 改為true
// step3: 找出map中 狀態還是處於false的城市,即為終點
func destCity(paths [][]string) string {
	hMap := map[string]bool{}

	for _, v := range paths {
		if _, ok := hMap[v[0]]; !ok {
			hMap[v[0]] = false
		}
		if _, ok := hMap[v[1]]; !ok {
			hMap[v[1]] = false
		}
	}
	for _, i := range paths {
		if _, ok := hMap[i[0]]; ok {
			hMap[i[0]] = true
		}
	}
	for v, k := range hMap {
		if k == false {
			return v
		}
	}
	return ""

}
