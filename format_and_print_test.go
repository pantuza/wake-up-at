package main

import (
	"bytes"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	baseTime := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)
	expOut := bytes.Buffer{}
	expOut.WriteString("To wake up at 8:00 AM, you should sleep at: 11:00 PM\n\n")
	expOut.WriteString("Also at: 12:30 AM | 2:00 AM | 3:30 AM\n")

	st := sleepTimes{
		sleepAt: []time.Time{
			baseTime.Add(-540 * time.Minute),
			baseTime.Add(-450 * time.Minute),
			baseTime.Add(-360 * time.Minute),
			baseTime.Add(-270 * time.Minute),
		},
	}

	output := st.format(baseTime)
	if output != expOut.String() {
		t.Fatalf("unexpected output string. \nGot:\n%v\nwant:\n%v", output,
			expOut.String())
	}
}
