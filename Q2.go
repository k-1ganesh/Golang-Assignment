/* 2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
constraint
1. Use const PI from the package math package
2. Use the Pow function from the math package
3. Round area float value to 2 decimal places
*/

package main

import (
	"math"
	"fmt"
	
)

func main() {
	fmt.Println("Enter the Value of Radius")
    var radius float64 
	fmt.Scanln(&radius)

	const PI = math.Pi
	area := PI * math.Pow(radius,2) 
	fmt.Printf("Area of Circle : %.2f" , area)
}