package main

import (
	"testing"
	"time"
)

func TestCalcTimes(t *testing.T) {
	baseTime := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

	st := sleepTimes{
		sleepAt: []time.Time{
			time.Time{},
			time.Time{},
			time.Time{},
			time.Time{},
		},
	}

	st.calcTimes(baseTime)

	if st.sleepAt[0].Hour() != 23 {
		t.Fatalf("First time shoulb be 23, got %d", st.sleepAt[0].Hour())
	} else if st.sleepAt[1].Hour() != 00 {
		t.Fatalf("Second time shoulb be 00, got %d", st.sleepAt[1].Hour())
	} else if st.sleepAt[2].Hour() != 2 {
		t.Fatalf("Third time shoulb be 2, got %d", st.sleepAt[2].Hour())
	} else if st.sleepAt[3].Hour() != 3 {
		t.Fatalf("Fourth time shoulb be 3, got %d", st.sleepAt[3].Hour())
	}
}
