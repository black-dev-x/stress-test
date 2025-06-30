package main

import (
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
	url, requests, concurrency := extractArguments()
	report := stress_tester.DoStressTest(url, requests, concurrency)
	report.Print()

}
