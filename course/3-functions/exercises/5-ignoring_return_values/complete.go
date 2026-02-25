package main

import "fmt"

func main1() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames1() (string, string) {
	return "John", "Doe"
}
