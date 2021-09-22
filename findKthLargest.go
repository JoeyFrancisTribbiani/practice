package main

func findKthLargest(nums []int, k int) int {
	// for _, v := range nums {
	// 	var p Node = Node{
	// 		val: v,
	// 		// next: nil
	// 	}
	// }
	return 0
}

type BigIntHeap struct {
	List     []int
	Capacity int
	Count    int
}

func (h *BigIntHeap) Init(nums []int, capacity int) {
	h.List = make([]int, capacity)
	h.Capacity = capacity
	h.Count = 0
	for _, v := range nums {
		h.Push(v)
	}
	// shishi
}

func (h *BigIntHeap) Swap(i, j int) {
	h.List[i], h.List[j] = h.List[j], h.List[i]
}
func (h *BigIntHeap) Pop() int {
	return h.List[0]
}
func (h *BigIntHeap) Push(num int) {
	if h.Len() == 0 {
		h.List[0] = num
		h.Count++
		return
	}
	tail := h.Len() - 1
	h.List[tail] = num
	p := tail / 2
	for p >= 0 {
		if h.List[tail] < h.List[p] {
			h.Swap(tail, p)
			tail = p
			p = tail / 2
		} else {
			break
		}
	}
	h.Count++
	if h.Count > h.Capacity {
		h.Count = h.Capacity
	}
}

func (h *BigIntHeap) Len() int {
	return h.Count
}
