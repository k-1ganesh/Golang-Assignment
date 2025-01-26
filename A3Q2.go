/*2. Word Count
Write a Program to fulfil the following conditions.
Input: A string containing words separated by space
Output: A slice containing words with the highest frequency in the input string.
Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
Condition: Words are case-sensitive. Joe is not the same as joe.
Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.
Example Test Case 1:
Input: My name is Joe and My Father's name is also Joe
Output: [My name is Joe]
Here, the words "My", "name", "is", "Joe" appeared 2 times each, which is also the highest frequency of any word.
Hence the output contains all 4 words.
Example Test Case 2:
Input: Europe is far but the US is farther
Output: [is]
Here, the word "is" appeared twice which is also the highest frequency of any word
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading string")
		return
	}

	input = strings.TrimSpace(input)
	values := strings.Split(input, " ")
	var mp = make(map[string]int)
	for _, value := range values {
		mp[value]++
	}
	max := 0
	for _, value := range mp {
		if value > max {
			max = value
		}
	}
	ans := []string{}
	for _, value := range values {
		if mp[value] == max {
			ans = append(ans, value)
			mp[value] = 0 // This is done to make sure that string is considered only once.
		}
	}
	fmt.Println(ans)
}
