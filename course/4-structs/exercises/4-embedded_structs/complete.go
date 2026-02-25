package main

import "fmt"

type sender1 struct {
	rateLimit int
	user1
}

type user1 struct {
	name   string
	number int
}

// don't edit below this line

func test1(s sender1) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main1() {
	test1(sender1{
		rateLimit: 10000,
		user1: user1{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test1(sender1{
		rateLimit: 5000,
		user1: user1{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test1(sender1{
		rateLimit: 1000,
		user1: user1{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
