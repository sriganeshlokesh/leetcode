## Plan of Attack

the idea behind this solution is to start filling in the `nums1` array from the end since there are empty values of length `n`.

## Algorithm

1. Initialize a variable `last` to hold the value of the last index of `nums1`.
2. Let `m` be the length of `nums1` and `n` be the length of `nums2`, as provided in the problem statement.
3. Loop through both lists until one of them finishes processing all its elements.
   **Note:** Since we are filling the values from the end of `nums1`, we fill in the greater value.
4. Check if the value from `nums1` is greater than the value from `nums2` at the respective ends of both lists.
5. If the value from `nums1` is greater, append it to the end of `nums1` and decrement `m` by 1.
   - Otherwise, append the value from `nums2` to the end of `nums1` and decrement `n` by 1.
6. Finally, after processing all the elements, check if there are any remaining elements in `nums2` and add them to the beginning of `nums1`.

### Time Complexity

The first while loop runs as long as both `m` and `n` are greater than 0, where `m` and `n` represent the lengths of `nums1` and `nums2` respectively.
Inside the loop, you compare elements from both arrays and place the larger element in `nums1`. This comparison and assignment step takes constant time `(O(1))`.
After each comparison and assignment, you decrement either `m` or `n`, depending on which element was larger.
The loop continues until either `m` or `n` becomes 0.
The second while loop runs when `n` is greater than 0 and copies the remaining elements from `nums2` to `nums1`.
In total, the time complexity of this code is `O(m + n)`, where `m` and `n` are the lengths of `nums1` and `nums2`, respectively. This is because you iterate through both arrays exactly once.

### Space Complexity

The code performs the merge in-place, which means it does not allocate any additional data structures that scale with the input size. It only uses a few extra variables (`last`, `m`, `n`) that do not depend on the size of the input arrays. As a result, the space complexity is `O(1)`, indicating constant space complexity.

### Summary

Time Complexity: `O(m + n)`
Space Complexity: `O(1)`
