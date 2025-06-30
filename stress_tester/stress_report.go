package stress_tester

import (
	"os"
	"os/exec"
)

type Report struct {
	totalTime       int64
	totalRequests   int
	successRequests int
	statusCodes     map[int]int
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func (r *Report) Print() {
	// clear()
	println("Total Time:", r.totalTime, "seconds")
	println("Total Requests:", r.totalRequests)
	println("Successful Requests:", r.successRequests)
	println("Processing Time:", r.totalTime)
	println("Status Codes:")
	for code, count := range r.statusCodes {
		println("  ", code, ":", count)
	}
}
