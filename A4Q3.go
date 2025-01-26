/* 3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

"Area :  <value>"
"Perimeter :  <value>"


HINT : Step 2 means, to define methods in Quadrilateral interface for given shapes


Formulae:

Area of Rectangle: length * width
Perimeter of Rectangle: 2*(length + breadth)


Area of Square: side * side
Perimeter of Square: 4 * side
*/

package main

import "fmt"

type quadrilateral interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length, breadth float64
}
type square struct {
	side float64
}

func (r rectangle) area() float64 {
	return r.breadth * r.length
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}

func print(q quadrilateral) {
	fmt.Printf("Area : %v\n", q.area())
	fmt.Printf("Perimeter: %v\n", q.perimeter())
}

func main() {
	var choice int
	fmt.Println("Enter the choice 1:Rectangle , 2:Square")
	fmt.Scanln(&choice)

	if choice == 1 {
		rect := rectangle{length: 10, breadth: 15}
		print(rect)
		return
	}
	if choice == 2 {
		sq := square{side: 15}
		print(sq)
		return
	}
	fmt.Println("Enter the correct choice")

}
