package main

import (
	"testing"
	"time"
)

func TestCalcTimes(t *testing.T) {

	baseTime := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

	var firstTime, secondTime, thirdTime, fourthTime time.Time
	calcTimes(&baseTime, &firstTime, &secondTime, &thirdTime, &fourthTime)

	if firstTime.Hour() != 23 {
		t.Fatalf("First time shoulb be 23, got %d", firstTime.Hour())
	} else if secondTime.Hour() != 00 {
		t.Fatalf("Second time shoulb be 00, got %d", secondTime.Hour())
	} else if thirdTime.Hour() != 2 {
		t.Fatalf("Third time shoulb be 2, got %d", thirdTime.Hour())
	} else if fourthTime.Hour() != 3 {
		t.Fatalf("Fourth time shoulb be 3, got %d", fourthTime.Hour())
	}
}
