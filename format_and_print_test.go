package main

import (
	"testing"
	"time"
)

func TestFormatAndPrint(t *testing.T) {

	baseTime := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

	time := Times{
		firstTime:  baseTime.Add(-540 * time.Minute),
		secondTime: baseTime.Add(-450 * time.Minute),
		thirdTime:  baseTime.Add(-360 * time.Minute),
		fourthTime: baseTime.Add(-270 * time.Minute),
	}
	time.formatAndPrint(&baseTime)
	// Output:
	// To wake up at 8:00 AM, you should sleep at: 11:00 PM
	//
	// Also at: 12:30 AM | 2:00 AM | 3:30 AM
}
