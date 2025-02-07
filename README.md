# Golang-Assignment
```go
package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	go func() {
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
```