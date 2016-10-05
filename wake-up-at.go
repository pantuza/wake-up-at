package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"time"
)

const morning string = "AM"
const evening string = "PM"

const midday int = 12
const timeFormat = "3:04 PM"

type sleepTimes struct {
	sleepAt []time.Time
}

func newSleepTimes(count int) *sleepTimes {
	st := sleepTimes{
		sleepAt: []time.Time{},
	}

	for i := 0; i < count; i++ {
		st.sleepAt = append(st.sleepAt, time.Time{})
	}
	return &st
}

// Parses the command line options
func parseOptions(hours, minutes *int, period *string) {
	flagHours := flag.Int("h", 8, "Hour to wake up")
	flagMinutes := flag.Int("m", 0, "Minute to wake up in the given hour")
	flagPeriod := flag.String("p", "am", "Period of the day: am/pm")

	flag.Parse()

	*hours = *flagHours
	*minutes = *flagMinutes
	*period = strings.ToUpper(*flagPeriod)

	if *period != morning && *period != evening {
		*period = morning
	}
}

// Returns true if the period of the day is Morning. Otherwise, false
func isMorning(period string) bool { return period == morning }

// Calculates the possibles times to go to sleep
func (st *sleepTimes) calcTimes(wakeTime time.Time) {
	offset := -540

	for i := range st.sleepAt {
		st.sleepAt[i] = wakeTime.Add(time.Duration(offset) * time.Minute)
		offset += 90
	}
}

// Formats the output message and print
func (st *sleepTimes) format(wakeTime time.Time) string {
	output := bytes.Buffer{}
	output.WriteString(fmt.Sprintf("To wake up at %v, ",
		wakeTime.Format(timeFormat)))

	for i := range st.sleepAt {
		if i == 0 {
			output.WriteString(fmt.Sprintf("you should sleep at: %v\n\n",
				st.sleepAt[i].Format(timeFormat)))
		} else if i == 1 {
			output.WriteString(fmt.Sprintf("Also at: %v",
				st.sleepAt[i].Format(timeFormat)))
		} else {
			output.WriteString(fmt.Sprintf(" | %v",
				st.sleepAt[i].Format(timeFormat)))
		}
	}

	return output.String()
}

func main() {
	var hours, minutes int
	var period string
	wakeTime := time.Now()
	st := newSleepTimes(4)

	// Parses command line options
	parseOptions(&hours, &minutes, &period)

	// If morning period, increment wake up day to tomorrow, else normalize
	// location that uses hours like 22 instead of 10pm
	if isMorning(period) {
		if wakeTime.Hour() > midday {
			wakeTime.AddDate(0, 0, 1)
		}
	} else {
		if hours <= midday {
			hours += midday
		}
	}

	// Update the wake time with correct values of hours and period
	wakeTime = time.Date(wakeTime.Year(), wakeTime.Month(),
		wakeTime.Day(), hours, minutes, 0, 0, wakeTime.Location())

	// Calculate possible times
	st.calcTimes(wakeTime)

	// Prints times to go sleep
	output := st.format(wakeTime)
	fmt.Println(output)
}
