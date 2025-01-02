/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

// Effiency improves with a two pass hash map, as demonstrated in https://leetcode.com/problems/two-sum/solutions/5486440/o-n-2-o-n-log-n-o-n-beats-100-java-python-c-javascript-go-rust/

/*
Brute Force - Nested Loops
*/
func twoSum(nums []int, target int) []int {
    var values []int

    for pointer1 := 0; pointer1 < len(nums); pointer1++ {
        fmt.Println("pointer1", pointer1)

        value1 := nums[pointer1]
        fmt.Println("value1", value1)

        for pointer2 := pointer1+1; pointer2 < len(nums); pointer2++ {
            fmt.Println("pointer2", pointer2)

            value2 := nums[pointer2]
            fmt.Println("value2", value2)

            if value1 + value2 == target {
                values = append(values, pointer1)
                values = append(values, pointer2)
                break
            }
        }

        if len(values) > 0 { 
            break
        }
    }

    return values
}
