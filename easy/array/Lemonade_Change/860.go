package Lemonade_Change

// Input: bills = [5,5,10,10,20]
// Output: false
// Explanation:
// From the first two customers in order, we collect two $5 bills.
// For the next two customers in order, we collect a $10 bill and give back a $5 bill.
// For the last customer, we can not give the change of $15 back because we only have two $10 bills.
// Since not every customer received the correct change, the answer is false.
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	if bills[0] == 10 || bills[0] == 20 {
		return false
	}
	for i := 0; i < len(bills); i++ {
		// 五塊 不用找零
		if bills[i] == 5 {
			five++
		}
		// 10塊 要找 5塊
		if bills[i] == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		}
		// 20塊 找1個10塊,1個五塊 || 找3個五塊
		if bills[i] == 20 {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
