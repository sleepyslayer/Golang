package main

import (
	"fmt"
	"time"
)

func sendMessage1(msg message1) {
	fmt.Println(msg.getMessage1())
}

type message1 interface {
	getMessage1() string
}

// don't edit below this line

type birthdayMessage1 struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage1) getMessage1() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport1 struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport1) getMessage1() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test1(m message1) {
	sendMessage1(m)
	fmt.Println("====================================")
}

func main1() {
	test1(sendingReport1{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test1(birthdayMessage1{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test1(sendingReport1{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test1(birthdayMessage1{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}
