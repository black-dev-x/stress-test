package main

import (
	"github.com/black-dev-x/stress-test/stress_tester"
)

func main() {

	url := "http://localhost:8080"
	requests := 1000
	concurrency := 10

	report := stress_tester.DoStressTest(url, requests, concurrency)
	report.Print()
}
