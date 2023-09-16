from typing import List


def remove_duplicates_from_sorted_array(nums: List[int]) -> int:
    slow = 1

    for fast in range(1, len(nums)):
        if nums[fast] != nums[fast - 1]:
            nums[slow] = nums[fast]
            slow += 1

    return slow


if __name__ == "__main__":
    print(remove_duplicates_from_sorted_array([1, 1, 2]))
    print(remove_duplicates_from_sorted_array([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]))
