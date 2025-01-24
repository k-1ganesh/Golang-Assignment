/* Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000


For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
Roman numerals are usually written from largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
Example 1:
Input: s = "III"
Output: 3
Explanation: III = 3.


Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.


Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

package main 
import ("fmt") 

func main() {
	var s string 
	fmt.Scanln(&s)
	var ans int = 0
	for i:=0 ; i<len(s);i++ {
        if s[i] == 'M' {
			ans += 1000
		} else if s[i] == 'D' {
			ans += 500
		} else if s[i] == 'C' {
			ans += 100
		} else if s[i] == 'L' {
			ans += 50
		} else if s[i] == 'X' {
			ans += 10
		} else if s[i] == 'V' {
			ans += 5
		} else {
			ans += 1
		}
	}
	for i:=0;i<len(s)-1;i++ {
		if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
			ans -= 200
		} else if (s[i] == 'L' && (s[i+1] == 'C' || s[i+1] == 'D')) {
			ans -= 100
		} else if (s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) {
			ans -= 20
		} else if (s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) {
			ans -= 2
		} else {

		}
	}

	fmt.Println(ans)

}