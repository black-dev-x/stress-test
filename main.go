package main

import (
	"os"
	"strconv"

	"github.com/black-dev-x/stress-test/stress_tester"
)

func extractArguments() (string, int, int) {
	var url string
	var requests int
	var concurrency int

	arguments := os.Args
	if len(arguments) < 7 {
		panic("Usage: stress-test --url: <URL> --requests: <number> --concurrency: <number>")
	}
	for i, arg := range arguments {
		switch arg {
		case "--url":
		case "--url:":
			url = arguments[i+1]
		case "--requests":
		case "--requests:":
			requests, _ = strconv.Atoi(arguments[i+1])
		case "--concurrency":
		case "--concurrency:":
			concurrency, _ = strconv.Atoi(arguments[i+1])
		}
	}

	return url, requests, concurrency
}

func main() {

	url, requests, concurrency := extractArguments()
	report := stress_tester.DoStressTest(url, requests, concurrency)
	report.Print()
}
