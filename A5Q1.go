/* 1. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value at that index in slice>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.

Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
Also, Print the following after handling panic

"internal error: <recovered panic value>"
*/

package main

import (
	"fmt"
)

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Internal error: %v\n", r)
		}
	}()

	if index >= 0 && index < len(slice) {
		fmt.Printf("item %v , value %v", index, slice[index])
		return
	}
	panic("Index out of range")

}

func main() {
	slice := []int{1, 2, 3, 4, 4, 5, 67, 23}
	accessSlice(slice, 15)
	accessSlice(slice[1:5], 3)
}
