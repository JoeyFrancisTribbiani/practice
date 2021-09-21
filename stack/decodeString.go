package stack

import (
	"container/list"
	"unicode"
)

func decodeString(s string) string {
	len := len(s)
	nStack := list.New()
	sStack := list.New()
	for i := 0; i < len; i++ {

		if s[i] == '[' {
			str := []byte{}
			for unicode.IsLetter(rune(s[i+1])) {
				str = append(str, s[i+1])
				i++
			}
			sStack.PushBack(str)
		}
	}
}

func decodeSub(s string) string {
	subString := ""
	if unicode.IsDigit(rune(s[i])) {
		n := int(s[i])
		nStack.PushBack(n)
	}
	str := ""
	for i := 0; i < count; i++ {
		str += subString
	}
	return str
}
