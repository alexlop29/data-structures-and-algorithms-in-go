/*
A pangram is a sentence where every letter of the English alphabet appears at least once.
Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

Example 1:
Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
Output: true
Explanation: sentence contains at least one of every letter of the English alphabet.

Example 2:
Input: sentence = "leetcode"
Output: false
 
Constraints:
1 <= sentence.length <= 1000
sentence consists of lowercase English letters.
*/

/*
Alogorithm:
- if string length is less than 26, return false
- Create slice of letters in alphabet
- Iterate through alphabet slice
- If the input does not contain an alphabet char, return false
- return true
*/
import (
    "strings"
)

func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false 
    }

    alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v","w", "x", "y", "z"}

    for _, v := range alphabet {
        if !strings.Contains(sentence, v){
            return false
        }
    }

    return true
}
