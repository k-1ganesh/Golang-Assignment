/* 1. The given program accepts an integer value between 1 to 4 from the user.
There is an option associated with each value, which is basically objects of different data types with some associated value.

Write a  method named "AcceptAnything" which takes any type of data as input.

Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

integer :
"This is a value of type Integer, <value>"

string :
"This is a value of type String, <value>"

boolean :
"This is a value of type Boolean, <value>"

Custom data type Hello :
"This is a value of type Hello, <value>"
*/

package main

import "fmt"

// Empty interface can have any value.
func AcceptAnything(val interface{}) {

	switch v := val.(type) {
	case int:
		fmt.Printf("This is the value of type Integer , %v", v)
	case string:
		fmt.Printf("This is the value of type String , %v", v)
	case bool:
		fmt.Printf("This is the value of type Boolean , %v", v)
	case Hello:
		fmt.Printf("This is the value of type Hello , %v", v)
	default:
		fmt.Println("Value of Unknown type is passed.")
	}

}

type Hello struct {
	name string
}

func main() {
	fmt.Println("Enter the input 1 to 4")
	var input int
	fmt.Scanln(&input)

	if input == 1 {
		AcceptAnything(10)
		return
	} 
	if input == 2 {
		AcceptAnything("Ganesh")
		return
	}
	if input == 3 {
		AcceptAnything(true)
		return
	}
	if input == 4 {
		AcceptAnything(Hello{name: "Ganesh"})
		return
	}
	fmt.Println("Enter valid input")
	
}
