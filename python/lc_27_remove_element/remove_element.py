from typing import List


def remove_element(nums: List[int], val: int) -> int:
    slow = 0

    for fast in range(len(nums)):
        if nums[fast] != val:
            nums[slow] = nums[fast]
            slow += 1

    return slow


if __name__ == "__main__":
    print(remove_element([3, 2, 2, 3], 3))
    print(remove_element([0, 1, 2, 2, 3, 0, 4, 2], 2))
