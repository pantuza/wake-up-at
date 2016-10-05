package main

import (
	"testing"
	"time"
)

func TestCalcTimes(t *testing.T) {

	baseTime := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

	var time Times
	time.calcTimes(&baseTime)

	if time.firstTime.Hour() != 23 {
		t.Fatalf("First time shoulb be 23, got %d", time.firstTime.Hour())
	} else if time.secondTime.Hour() != 00 {
		t.Fatalf("Second time shoulb be 00, got %d", time.secondTime.Hour())
	} else if time.thirdTime.Hour() != 2 {
		t.Fatalf("Third time shoulb be 2, got %d", time.thirdTime.Hour())
	} else if time.fourthTime.Hour() != 3 {
		t.Fatalf("Fourth time shoulb be 3, got %d", time.fourthTime.Hour())
	}
}
