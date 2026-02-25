package main

import "fmt"

type authenticationInfo1 struct {
	username string
	password string
}

func (authInfo authenticationInfo1) getBasicAuth() string {
	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
}

// don't touch below this line

func test1(authInfo authenticationInfo1) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main1() {
	test1(authenticationInfo1{
		username: "Google",
		password: "12345",
	})
	test1(authenticationInfo1{
		username: "Bing",
		password: "98765",
	})
	test1(authenticationInfo1{
		username: "DDG",
		password: "76921",
	})
}
