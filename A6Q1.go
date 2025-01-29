/* 1. Given a string containing conversation between alice and bob. In the string, if it reaches $, it means end of alice message and if it reaches #, it means end of bob's message. and if it reaches ^,
it means end of conversation ignore string after that.

Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example.

write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?
*/

package main

import (
	"fmt"
)

func helper(input string, aliceChannel chan string, bobChannel chan string, flag chan bool) {
	defer close(aliceChannel)
	defer close(bobChannel)
	defer close(flag)
	var msg string = ""
	for _, value := range input {
		if value == '^' {
			flag <- true
			return
		} else if value == '$' {
			aliceChannel <- msg
			msg = ""
		} else if value == '#' {
			bobChannel <- msg
			msg = ""
		} else {
			msg = msg + string(value)
		}
	}
}

func main() {
	var input string
	fmt.Scanln(&input)
	aliceChannel := make(chan string)
	bobChannel := make(chan string)
	flag := make(chan bool)
	go helper(input, aliceChannel, bobChannel, flag)
	firstMsg := true
	for {
		select {
		case val := <-aliceChannel:
			if !firstMsg {
				fmt.Printf(",")
			}
			fmt.Printf("alice : %v", val)
			firstMsg = false
		case val := <-bobChannel:
			if !firstMsg {
				fmt.Printf(",")
			}
			fmt.Printf("bob : %v", val)
			firstMsg = false
		case <-flag:
			return
		}
	}
}
