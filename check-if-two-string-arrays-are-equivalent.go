/*
Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.
A string is represented by an array if the array elements concatenated in order forms the string.

Example 1:
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.

Example 2:
Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false

Example 3:
Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true

Constraints:
1 <= word1.length, word2.length <= 103
1 <= word1[i].length, word2[i].length <= 103
1 <= sum(word1[i].length), sum(word2[i].length) <= 103
word1[i] and word2[i] consist of lowercase letters.
*/

/*
Algorithm:
- create placeholder values to hold the concatenated strings for each array
- iterate through each array and concacate the current word to the strin
- compare
- return true or false
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    str1 := ""
    str2 := ""

    for _, v := range word1 {
        str1 = str1 + v
    }

    for _, v := range word2 {
        str2 = str2 + v
    }

    return str1 == str2
}
