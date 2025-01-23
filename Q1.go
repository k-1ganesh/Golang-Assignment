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
	input,err := reader.ReadString('\n')
	if err!=nil{
		fmt.Errorf("error in reading input")
		return
	}
	
	input = strings.TrimSpace(input)

	values := strings.Split(input , ",")

	if len(values) != 3 {
		fmt.Println("Enter the valid input")
		return
	}

	Principal,err:= strconv.ParseFloat(values[0] , 64)
	if err != nil {
		fmt.Errorf("Not a Number")
		return
	}

	Rate,err := strconv.ParseFloat(values[1] , 64)
	if err != nil {
		fmt.Errorf("Not a Number")
		return
	}

	Time,err:= strconv.ParseFloat(values[2] , 64)
	if err != nil {
		fmt.Errorf("Not a Number")
		return
	}

	Compound_Interest := Principal * Rate * Time / 100 
	fmt.Printf("Compound Interest is: %.2f" , Compound_Interest)
}