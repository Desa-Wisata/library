package utils

import "time"

// TimeNow ...
func TimeNow() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	start := time.Now().In(loc)

	return start
}

// TimeDuration ...
func TimeDuration(t time.Time) float64 {
	duration := time.Since(t).Seconds()

	return duration
}
