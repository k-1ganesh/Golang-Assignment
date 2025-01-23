/*1. Write a program to calculate the simple interest
First-line has the comma-separated values of the Principal, rate and time (in years) respective
*constraints: Round simple interest float value to 2 decimal places
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value of Principle,Rate,Time")
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	values := strings.Split(input , ",")

	Principal,_ := strconv.ParseFloat(values[0] , 64)
	Rate,_ := strconv.ParseFloat(values[1] , 64)
	Time,_ := strconv.ParseFloat(values[2] , 64)

	Compound_Interest := Principal * Rate * Time / 100 
	fmt.Printf("Compound Interest is: %.2f" , Compound_Interest)
}