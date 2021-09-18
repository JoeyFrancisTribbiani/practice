package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1 := m - 1
	n2 := n - 1
	k := m + n - 1
	for i := k; i >= 0; i-- {
		if n1 >= 0 {
			if n2 < 0 {
				break
			}
			if nums1[n1] < nums2[n2] {
				nums1[i] = nums2[n2]
				n2--
			} else {
				nums1[i] = nums1[n1]
				n1--
			}
		} else {
			nums1[i] = nums2[n2]
			n2--
		}
	}
}
