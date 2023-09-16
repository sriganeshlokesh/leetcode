## Plan of Attack

the idea behind this solution is to keep two pointers `slow` and `fast` which iterate through the array and weed out any items which match the value of the **target**. the `slow` variable will keep track of all the non-target elements and `fast` will iterate over the entire array.

## Algorithm

1. initialize the `slow` variable to 0, which represents the current position for the next non-target element.
2. iterate through the array using the `fast` pointer.
3. for each element, check if it is not equal to the target
   i. if true, then replace the element at the `slow` pointer with the element at the fast pointer.
   ii. increment the `slow` pointer by 1.
4. perform the steps until we have processed all the elements in the array.
5. return the value of `slow`, which represents the length of the modified array without the **target** value

### Time Complexity

The algorithm uses a two-pointer approach with a `slow` and a `fast` pointer. It iterates through the nums list once. In the worst case, when there are no occurrences of **target** in the list, it will iterate through the entire list of length `n`. Therefore, the time complexity of this algorithm is `O(n)`, where n is the length of the nums list.

### Space Complexity

The algorithm modifies the input list nums in-place and doesn't use any additional data structures that grow with the input size. It only uses two integer variables, `slow` and `fast`, to keep track of positions within the list. Therefore, the space complexity of this algorithm is `O(1)`, indicating that it uses a constant amount of extra space regardless of the size of the input list.
