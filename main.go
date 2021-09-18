package main

import (
	"fmt"
	"practice/stack"
	"time"
)

func main() {
	start1 := time.Now()
	s := stack.Constructor()
	s.Push(2147483646)
	s.Push(2147483646)
	s.Push(2147483647)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Push(2147483647)
	s.Push(-2147483648)
	top := s.Top()
	elapsed1 := time.Since(start1)
	fmt.Println(top)
	fmt.Println("耗时:", elapsed1)
}
