package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	sub := []rune(s)
	hashmap := make(map[rune]int)

	var last int
	var longest int
	for i := 1; i < len(sub)+1; i++ {
		letter := sub[i-1]
		if index, ok := hashmap[letter]; ok && index > last {
			last = index
		} else {
			longest = max(longest, i-last)
		}

		hashmap[letter] = i
	}

	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
