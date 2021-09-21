package stack

import (
	"container/list"
	"strconv"
	"strings"
)

func EvalRPN(tokens []string) int {
	operators := "+-*/"
	stack := list.New()
	first := true
	sum := 0
	if len(tokens) == 1 {
		sum, _ = strconv.Atoi(tokens[0])
		return sum
	}
	for i := 0; i < len(tokens); i++ {
		if strings.Index(operators, tokens[i]) < 0 {
			stack.PushBack(tokens[i])
		} else {
			a, b := 0, 0
			if first {
				e1 := stack.Back()
				stack.Remove(e1)
				b, _ = strconv.Atoi(e1.Value.(string))
				e2 := stack.Back()
				stack.Remove(e2)
				a, _ = strconv.Atoi(e2.Value.(string))
				first = false
			} else {
				if strings.Index(operators, tokens[i-1]) < 0 {
					a = sum
					e1 := stack.Back()
					stack.Remove(e1)
					b, _ = strconv.Atoi(e1.Value.(string))
				} else {
					b = sum
					e1 := stack.Back()
					stack.Remove(e1)
					a, _ = strconv.Atoi(e1.Value.(string))

				}
			}
			if tokens[i] == "+" {
				sum = a + b
			}
			if tokens[i] == "-" {
				sum = a - b
			}
			if tokens[i] == "*" {
				sum = a * b
			}
			if tokens[i] == "/" {
				sum = a / b
			}
		}
	}
	return sum
}
