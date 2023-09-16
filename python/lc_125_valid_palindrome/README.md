## Plan of Attack

the idea behind this solution is to use two pointer `i` and `j` where i starts from the beginning of the string and j starts from the end of the string.
we also need to account for alphanumeric characters and ignore them

## Algorithm

1. initialize two pointers and move them from opposite ends.
2. the first pointer `i` starts at the beginning of the string and move towards the middle, while second pointer `j` starts at the end and moves toward the middle.
3. compare the elements at each position to detect a non-matching pair
   1. ignore any alphanumeric characters while comparing the indices.
4. if both the pointers reach the middle of the string without encountering a non-matching pair, the string is a palindrome.

### Time Complexity

The code uses two pointers `i` and `j` and iterates through the string `s` from both ends towards the center.
Inside the while loop, it skips non-alphanumeric characters by advancing `i` and `j` accordingly. This operation takes `O(n)` time in the worst case, where `n` is the length of the string.
Then, it compares characters at the `i-th` and `j-th` positions (after converting them to lowercase). This comparison takes constant time `O(1)`.
Therefore, the overall time complexity of this code is `O(n)`, where `n` is the length of the input string `s`.

### Space Complexity

The code uses a constant amount of extra space regardless of the size of the input string. It only uses two pointers (`i` and `j`) and a few variables for character comparisons and loop control.
The space complexity is `O(1)`, constant space complexity.

### Summary

Time Complexity: `O(n)`
Space Complexity: `O(1)`
