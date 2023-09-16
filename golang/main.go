package main

import (
	"fmt"

	lc125validpalindrome "github.com/sriganeshlokesh/leetcode/internal/lc_125_valid_palindrome"
	lc26removeduplicatesfromsortedarray "github.com/sriganeshlokesh/leetcode/internal/lc_26_remove_duplicates_from_sorted_array"
	lc27removeelement "github.com/sriganeshlokesh/leetcode/internal/lc_27_remove_element"
	lc88mergesortedarrays "github.com/sriganeshlokesh/leetcode/internal/lc_88_merge_sorted_arrays"
)

func main() {
	// LC - 26 - Remove Duplicates From Sorted Array
	fmt.Println(lc26removeduplicatesfromsortedarray.RemoveDuplicatesFromSortedArray([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	// LC - 27 - Remove Element
	fmt.Println(lc27removeelement.RemoveElement([]int{3, 2, 2, 3}, 3))
	// LC - 88 - Merge Sorted Arrays
	fmt.Println(lc88mergesortedarrays.MergeSortedArrays([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	// LC - 125 - Valid Palindrome
	fmt.Println(lc125validpalindrome.IsPalindrome("A man, a plan, a canal: Panama"))
}
