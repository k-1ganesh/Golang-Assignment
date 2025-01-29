/* 2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched

e.g: Input: test123 output: 321tset 2
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func reverse(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	strByte := []byte(str)
	left := 0
	right := len(strByte) - 1

	for left < right {
		strByte[left], strByte[right] = strByte[right], strByte[left]
		left++
		right--
	}

	str = string(strByte)
	fmt.Println(str, runtime.NumGoroutine())

}

func main() {
	input := "test123"
	var wg sync.WaitGroup

	wg.Add(1)
	go reverse(input, &wg)
	wg.Wait()
}
