package stack

import (
	"container/list"
	"strings"
)

func isValid(s string) bool {
	// chars := []byte{'(')'{'}'[']'}
	chars := "(){}[]"
	stack := list.New()
	for i := 0; i < len(s); i++ {
		k := strings.IndexByte(chars, s[i])
		if k%2 == 0 {
			stack.PushBack(s[i])
		} else {
			e := stack.Back()
			if e == nil {
				return false
			}
			if strings.IndexByte(chars, e.Value.(byte)) != k-1 {
				return false
			}
			stack.Remove(e)
		}
	}
	e := stack.Back()
	if e != nil {
		return false
	}
	return true
}
