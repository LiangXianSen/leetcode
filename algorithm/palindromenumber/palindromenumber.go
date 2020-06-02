package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	list := splitNum(x)
	length := len(list)
	if length == 1 || length == 0 { // based on splitNum, only one number should be true
		return true
	}

	for i := 1; i < length+1/2; i++ {
		if list[i-1] != list[length-i] {
			return false
		}
	}
	return true
}

func splitNum(x int) (list []int) {
	for x != 0 {
		list = append(list, x%10)
		x /= 10
	}
	return
}
