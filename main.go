package main

import (
	"fmt"
	"practice/set"
)

func main() {
	// start1 := time.Now()
	// s := stack.Constructor()
	// s.Push(2147483646)
	// s.Push(2147483646)
	// s.Push(2147483647)
	// s.Pop()
	// s.Pop()
	// s.Pop()
	// s.Push(2147483647)
	// s.Push(-2147483648)
	// top := s.Top()
	// elapsed1 := time.Since(start1)
	// fmt.Println(top)
	// fmt.Println("耗时:", elapsed1)

	// rpns := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// rpns := []string{"3", "11", "+", "5", "-"}
	// sum := stack.EvalRPN(rpns)
	// fmt.Println(sum)

	// nums := []int{1}
	// nums := []int{1, 1, 1, 1, 1}
	// fmt.Println(stack.FindTargetSumWays(nums, 1))
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(set.ThreeSum(nums))
}
