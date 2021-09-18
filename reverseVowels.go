package main

import (
	"bytes"
)

func reverseVowels(s string) string {
	voweles := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	res := []rune(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < len(s)-1 && bytes.IndexByte(voweles, s[i]) == -1 {
			if i < j {
				i++
			} else {
				break
			}
		}
		for j > 0 && bytes.IndexByte(voweles, s[j]) == -1 {
			if j > i {
				j--
			} else {
				break
			}
		}
		if i < j {
			res[i], res[j] = res[j], res[i]
		}
		i++
		j--
	}
	return string(res)
}
