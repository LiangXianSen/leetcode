package removeelement

func removeElement(nums []int, val int) int {
	var i int
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
