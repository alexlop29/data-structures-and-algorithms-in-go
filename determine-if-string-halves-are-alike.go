/*
You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.
Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and lowercase letters.
Return true if a and b are alike. Otherwise, return false.

Example 1:
Input: s = "book"
Output: true
Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.

Example 2:
Input: s = "textbook"
Output: false
Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
Notice that the vowel o is counted twice.
 
Constraints:
2 <= s.length <= 1000
s.length is even.
s consists of uppercase and lowercase letters.
*/

/*
- Split string in half and assign each half to a var
- Iterate through each var and keep count of vowels
- compare vowel counts
- if equal, return true / if not equal, return false
*/

import (
    "slices"
)

func halvesAreAlike(s string) bool {
    length := len(s)

    firstHalfCount := 0
    secondHalfCount := 0

    vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

    for i, r := range s {
        if !slices.Contains(vowels, string(r)){
            continue
        }

        half := length / 2

        if i < half {
            firstHalfCount++
        } else {
            secondHalfCount++
        }
    }

    return firstHalfCount == secondHalfCount
}
