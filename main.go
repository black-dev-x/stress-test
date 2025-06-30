package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/black-dev-x/stress-test/stress_tester"
)

func extractArguments() (string, int, int) {

	arguments := strings.Join(os.Args[1:], " ")
	urlRegex := "url[:=\\s]*(\\S*)"
	requestsRegex := "requests[:=\\s]*(\\d*)"
	concurrencyRegex := "concurrency[:=\\s]*(\\d*)"

	urlMatch := regexp.MustCompile(urlRegex).FindStringSubmatch(arguments)
	requestsMatch := regexp.MustCompile(requestsRegex).FindStringSubmatch(arguments)
	concurrencyMatch := regexp.MustCompile(concurrencyRegex).FindStringSubmatch(arguments)

	var url string
	var requests int
	var concurrency int

	if len(urlMatch) > 1 {
		url = urlMatch[1]
	}
	if len(requestsMatch) > 1 {
		requests, _ = strconv.Atoi(requestsMatch[1])
	}
	if len(concurrencyMatch) > 1 {
		concurrency, _ = strconv.Atoi(concurrencyMatch[1])
	}

	return url, requests, concurrency
}

func main() {
	println("Extracting arguments...")
	url, requests, concurrency := extractArguments()
	println("URL:", url)
	println("Requests:", requests)
	println("Concurrency:", concurrency)

	println("Starting stress test...")
	report := stress_tester.DoStressTest(url, requests, concurrency)

	println("Printing last report...")
	report.Print()

	println("Stress test completed successfully.")
	fmt.Printf("%v", os.Args)
}
