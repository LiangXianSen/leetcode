package reverseinteger

func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if !(-1<<31 <= y && y <= 1<<31-1) {
			return 0
		}
		x /= 10
	}
	return y
}
