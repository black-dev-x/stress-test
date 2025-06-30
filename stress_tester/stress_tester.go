package stress_tester

import (
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	url         string
	requests    int
	concurrency int
}

func DoStressTest(url string, requests, concurrency int) *Report {
	startTime := time.Now().Unix()
	report := newReport()
	statusCodeChannel := make(chan int, concurrency)
	errorChannel := make(chan error, concurrency)
	for i := 0; i < requests; i++ {
		go func() {
			response, err := http.Get(url)
			if err != nil {
				errorChannel <- err
				return
			}
			defer response.Body.Close()
			statusCodeChannel <- response.StatusCode
		}()
	}

	for i := 0; i < requests; i++ {
		select {
		case statusCode := <-statusCodeChannel:
			report.statusCodes[statusCode]++
			report.totalRequests++
			if statusCode < 400 {
				report.successRequests++
			}
		case err := <-errorChannel:
			fmt.Println("Error occurred:", err)
		}
	}

	endTime := time.Now().Unix()
	report.totalTime = endTime - startTime
	return report
}

func newRequest(url string, requests, concurrency int) *Request {
	return &Request{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
	}
}

func newReport() *Report {
	return &Report{
		statusCodes: make(map[int]int),
	}
}
