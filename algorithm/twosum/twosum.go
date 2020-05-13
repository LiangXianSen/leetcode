package twosum

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for index, value := range nums {
		anotherNum := target - value
		if _, ok := hashmap[anotherNum]; ok {
			return []int{hashmap[anotherNum], index}
		}
		hashmap[value] = index
	}
	return nil
}

func twoSumWithLoop(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
