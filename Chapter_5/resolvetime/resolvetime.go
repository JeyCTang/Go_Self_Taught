package main

import "fmt"

const (
	// SecondsPerMinute define how many seconds in per minute
	SecondsPerMinute = 60
	// SecondsPerHour define how many seconds in per hour
	SecondsPerHour = SecondsPerMinute * 60
	// SecondsPerDay define how many seconds in per day
	SecondsPerDay = SecondsPerHour * 24
)

// parse second to minute, hour, and day
func resolveTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute

	return
}
func main() {
	// print returned values of resolveTime
	fmt.Println(resolveTime(10000))
	// only get hour and minute
	_, hour, minute := resolveTime(10000)
	fmt.Println(hour, minute)
	// only get day
	day, _, _ := resolveTime(10000)
	fmt.Println(day)
}
