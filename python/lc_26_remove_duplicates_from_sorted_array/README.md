## Plan of Attack

the idea is to use two pointers `slow` and `fast`, to iterate over the array of integers. the variable `slow` is used to keep track of the current index where a unique element should be inserted. we start with assigning `fast` value with 1 since the first element of the array is always unique.

## Algorithm

1. start looping through the array of integers from 1 to last element of the array.
2. check if the current element is not equal to the previous element to find the next unique element in the array.
3. replace the element in the `slow` index with the element in the `fast` index and increment the value of `slow` by 1.
4. once we follow this process until the end of the array, we would have effectively overwritten any duplicates up until the `slow` index.
5. the `slow` value represents the length of the array with duplicates removed. return `slow`.

### Time Complexity

The algorithm iterates through the input array once with two pointers, `slow` and `fast`. Since it traverses the entire array once, the time complexity is `O(n)`, where n is the length of the input array nums.

### Space Complexity

The algorithm modifies the input array in-place, which means it doesn't use any additional data structures that grow with the input size. It only uses two integer variables, `slow` and `fast`, to keep track of positions within the array. Therefore, the space complexity of this algorithm is `O(1)`, indicating that it uses a constant amount of extra space regardless of the size of the input array.
