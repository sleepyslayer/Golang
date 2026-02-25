package main

import "fmt"

func main1() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends1(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}
