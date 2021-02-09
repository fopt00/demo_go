package tour

import "strings"

func WordCount(text string) map[string]int {
	cnt := map[string]int{}
	fields := strings.Fields(text)
	for _, word := range fields {
		cnt[strings.ToLower(word)]++
	}
	return cnt
}
