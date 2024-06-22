package utils

import "time"

func GetCurrentDateEndTime() string {
	currentTime := time.Now()
	endOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	return endOfDay.Format("2006-01-02 15:04:05")
}

func GetCurrentDateStartTime() string {
	currentTime := time.Now()
	startOfDay := currentTime.Format("2006-01-02 00:00:00")
	return startOfDay
}
