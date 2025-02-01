/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1

Constraints:
1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

/*
Algorithm:
- Iterate through nums.
- For each num apperance, add it to the hash map.
- Iterate through hash, and return the value with a count of 1
*/

func singleNumber(nums []int) int {
    hashOfCounts := map[int]int{}

    for _, currVal := range nums {
        v, ok := hashOfCounts[currVal]
        if (ok){
            hashOfCounts[currVal] = v + 1
        } else {
            hashOfCounts[currVal] = 1
        }
    }

    unique := -1

    for k, v := range hashOfCounts {
        if v == 1 {
            unique = k
            break;
        }
    };

    return unique
}
