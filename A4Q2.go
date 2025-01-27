/* 2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

Hint:
Method Signatures for Area and Perimeter
Area() float64
Perimeter() float64

Formulae:
Area = length * width
Perimeter = 2 * (length + width)

Example Test Case 1 :
Input: 10 20
Output:
Area of Rectangle: 200
Perimeter of Rectangle: 60
*/

package main

import "fmt"

type quadrilateral interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.breadth * r.length
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func printInfo(q quadrilateral) {
	fmt.Printf("Area of rectangle is %v\n", q.area())
	fmt.Printf("Perimeter of rectangle is %v\n", q.perimeter())
}

func main() {
	fmt.Println("Enter the value of Lenght and Breadth")
	var length, breadth float64
	fmt.Scanln(&length, &breadth)
	rect := rectangle{length: length, breadth: breadth}
	printInfo(rect)
}
