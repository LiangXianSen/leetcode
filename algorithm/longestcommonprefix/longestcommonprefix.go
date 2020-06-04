package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	var prefix []string

	if len(strs) == 0 {
		return ""
	}

	prefix = strings.Split(strs[0], "")
	for i := 1; i < len(strs); i++ {
		if len(prefix) == 0 {
			break
		}
		for j := 0; ; j++ {
			// 比prefix短，裁剪prefix
			if j > len(strs[i])-1 {
				prefix = prefix[0:j]
				break
			}

			if j > len(prefix)-1 {
				break
			}

			// 不等于，当前索引裁剪
			if string(strs[i][j]) != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return strings.Join(prefix, "")
}
