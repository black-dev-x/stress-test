package stress_tester

type Report struct {
	totalTime       int64
	totalRequests   int
	successRequests int
	statusCodes     map[int]int
}

func (r *Report) Print() {
	println("Total Time:", r.totalTime, "seconds")
	println("Total Requests:", r.totalRequests)
	println("Successful Requests:", r.successRequests)

	println("Status Codes:")
	for code, count := range r.statusCodes {
		println("  ", code, ":", count)
	}
}
