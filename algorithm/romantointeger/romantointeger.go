package romantointeger

var hashmap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && hashmap[string(s[i])] < hashmap[string(s[i+1])] {
			sum -= hashmap[string(s[i])]
		} else {
			sum += hashmap[string(s[i])]
		}
	}
	return sum
}
