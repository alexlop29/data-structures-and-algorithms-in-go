/*
Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.
Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.

Example 1:
Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
Output: [2,2,2,1,4,3,3,9,6,7,19]

Example 2:
Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
Output: [22,28,8,6,17,44]

Constraints:
1 <= arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
All the elements of arr2 are distinct.
Each arr2[i] is in arr1.
*/

/*
Alogorithm:
- Create a slice to keep track of sorted elements
- Iterate through arr2
- For each element in arr2:
    - Maintain counter to keep track of appearances in arr1
    - Check for matching value in arr1
    - If found, aggregate counter and pop the element from arr1
    - Continue until the value is not found in arr1
    - If not found, append the count of the value to the sorted arr
- Return sorted arr
*/

import (
    "slices"
    "sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
    sortedArr := []int{}

    for _, v := range arr2 {
        for {
            index := slices.Index(arr1, v)
            if index == -1 {
                break
            }
            fmt.Println(v, index, arr1)
            arr1 = append(arr1[:index], arr1[index+1:]...)
            sortedArr = append(sortedArr, v)
        }
    }

    sort.Ints(arr1)
    sortedArr = append(sortedArr, arr1...)

    return sortedArr
}
