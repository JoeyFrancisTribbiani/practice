package set

func Intersect(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	res := []int{}
	for _, v := range nums1 {
		//c,ok:=hash[v]
		hash[v]++
	}
	for _, v1 := range nums2 {
		if hash[v1] > 0 {
			res = append(res, v1)
			hash[v1]--
		}
	}
	return res
}
