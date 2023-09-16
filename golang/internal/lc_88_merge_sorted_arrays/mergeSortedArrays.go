package lc88mergesortedarrays

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {

	last := m + n - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}

		last--
	}

	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}

	return nums1
}
