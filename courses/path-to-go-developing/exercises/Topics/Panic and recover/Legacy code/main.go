package main

import "fmt"

const (
	FiveTasks  = 5
	ThreeTasks = 3
	OneTask    = 1
)

func scanTasks() (int, int) {
	var active, pending int
	_, err := fmt.Scan(&active, &pending)
	if err != nil {
		return 0, 0
	}
	return active, pending
}

func getTotalTasks() int {
	activeTasks, pendingTasks := scanTasks()
	return activeTasks + pendingTasks
}

func main() {
	totalTaskCount := getTotalTasks()

	switch {
	case totalTaskCount >= FiveTasks:
		fmt.Println("overload")
	case totalTaskCount >= ThreeTasks:
		fmt.Println("average")
	case totalTaskCount >= OneTask:
		fmt.Println("low")
	default:
		fmt.Println("idle")
	}
}
