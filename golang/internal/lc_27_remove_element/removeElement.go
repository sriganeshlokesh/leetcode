package lc27removeelement

func RemoveElement(nums []int, target int) int {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
