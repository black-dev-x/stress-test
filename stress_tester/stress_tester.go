package stress_tester

import (
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
	limit := make(chan bool, concurrency)

	go func() {
		for i := 0; i < requests; i++ {
			limit <- true
			go func() {
				defer func() { <-limit }()
				response, err := http.Get(url)
				if err != nil {
					errorChannel <- err
					return
				}
				defer response.Body.Close()
				statusCodeChannel <- response.StatusCode
			}()
		}
	}()

	go func() {
		for {
			<-time.After(time.Second * 1)
			report.Print()
		}
	}()

	for i := 0; i < requests; i++ {
		endTime := time.Now().Unix()
		report.totalTime = endTime - startTime
		select {
		case statusCode := <-statusCodeChannel:
			report.statusCodes[statusCode]++
			report.totalRequests++
			if statusCode < 400 {
				report.successRequests++
			}
		case <-errorChannel:

		}
	}

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
