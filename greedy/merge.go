package greedy

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] > intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, v := range intervals {
		if len(merged) != 0 && v[0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1] = []int{merged[len(merged)-1][0], int(math.Max(float64(v[1]), float64(merged[len(merged)-1][1])))}
		} else {
			merged = append(merged, v)
		}
	}
	return merged
}
