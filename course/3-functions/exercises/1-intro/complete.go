package main

import "fmt"

func concat1(s1 string, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func main1() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func test1(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
