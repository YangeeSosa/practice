package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := append(nums1, nums2...)
	sort.Ints(mergedArr)

	n := len(mergedArr)

	if n%2 != 0 {
		return float64(mergedArr[n/2])
	} else {
		mid1 := mergedArr[n/2-1]
		mid2 := mergedArr[n/2]
		return float64(mid1+mid2) / 2.0
	}
}
