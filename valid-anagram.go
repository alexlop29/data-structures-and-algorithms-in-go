/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
*/

import "reflect"

func countAppearances(str string) map[string]int {
    appearances := make(map[string]int)

    for _, runeVal := range str {
        char := string(runeVal)
        value, ok := appearances[char]
        if ok {
            appearances[char] = value + 1
        } else {
            appearances[char] = 1
        }
    }

    return appearances
}

func isAnagram(s string, t string) bool {
    sMap := countAppearances(s)
    tMap := countAppearances(t)

    return reflect.DeepEqual(sMap, tMap)
}
