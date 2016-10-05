package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

const morning string = "AM"
const evening string = "PM"

const midday int = 12
const timeFormat = "3:04 PM"

type Times struct {
	firstTime  time.Time
	secondTime time.Time
	thirdTime  time.Time
	fourthTime time.Time
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
func (t *Times) calcTimes(wakeTime *time.Time) {

	t.firstTime = wakeTime.Add(-540 * time.Minute)
	t.secondTime = wakeTime.Add(-450 * time.Minute)
	t.thirdTime = wakeTime.Add(-360 * time.Minute)
	t.fourthTime = wakeTime.Add(-270 * time.Minute)
}

// Formats the output message and print
func (t *Times) formatAndPrint(wakeTime *time.Time) {

	fmt.Printf("To wake up at %v, ", wakeTime.Format(timeFormat))
	fmt.Printf("you should sleep at: %v\n\n", t.firstTime.Format(timeFormat))
	fmt.Printf("Also at: %v | %v | %v\n", t.secondTime.Format(timeFormat),
		t.thirdTime.Format(timeFormat), t.fourthTime.Format(timeFormat))
}

func main() {

	now := time.Now()

	// Variables for wake time calculation
	var period string
	var hours, minutes int

	// Parses command line options
	parseOptions(&hours, &minutes, &period)

	wakeTime := now

	if isMorning(period) {

		// Increment wake up day to tomorrow
		if now.Hour() > midday {
			wakeTime = wakeTime.AddDate(0, 0, 1)
		}

		// If wake up period is PM
	} else {

		// Normalize location that uses hours like 22 instead of 10 PM.
		if hours <= midday {
			// then we increment it by 12 hours
			hours += midday
		}
	}

	// Update the wake time with correct values of hours and period
	wakeTime = time.Date(wakeTime.Year(), wakeTime.Month(),
		wakeTime.Day(), hours, minutes, 0, 0, wakeTime.Location())

	// Calculate possible times
	var time Times
	time.calcTimes(&wakeTime)

	// Prints times to go sleep
	time.formatAndPrint(&wakeTime)
}
