package set

import "strings"

func frequencySort(s string) string {
	len := len([]rune(s))
	strs := make([]string, len)
	dict := map[string]int{}
	for i := range s {
		dict[string(s[i])]++
	}
	for k, v := range dict {
		strs[v] += strings.Repeat(k, v)
	}
	res := ""
	for i := len - 1; i > 0; i-- {
		if strs[i] != "" {
			res += strs[i]
		}
	}
	return res
}
