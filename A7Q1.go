// Reason why code was failing because of irregularities between the execution of both ananomous function.
// For consistent output the function incrementing n should be executed after the function checking n is odd or even.

package main

import (
	"fmt"
	"time"
	"sync"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
    var wg sync.WaitGroup
	go func() {
		defer wg.Done() 
		wg.Add(1)
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		wg.Wait() // This blocks function until above function completes its execution.
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}