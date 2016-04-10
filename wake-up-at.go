

package main


import (
    "fmt"
    "time"
    "flag"
    "strings"
)


const MORNING string = "AM"
const EVENING string = "PM"


const MIDDAY int = 12


// Parses the command line options
func parseOptions(hours, minutes *int, period *string) {

    flagHours := flag.Int("h", 8, "Hour to wake up")
    flagMinutes := flag.Int("m", 0, "Minute to wake up in the given hour")
    flagPeriod := flag.String("p", "am", "Period of the day: am/pm")

    flag.Parse()

    *hours = *flagHours
    *minutes = *flagMinutes
    *period = strings.ToUpper(*flagPeriod)

    if *period != MORNING && *period != EVENING {
        *period = MORNING
    }
}


// Returns true if the period of the day is Morning. Otherwise, false
func isMorning (period string) bool { return period == MORNING }


func main () {

    now := time.Now()

    // Variables for wake time calculation
    var period string
    var hours int
    var minutes int


    // Parses command line options
    parseOptions(&hours, &minutes, &period)


    wake_time := now

    if isMorning(period) {

        if now.Hour() > MIDDAY {
            wake_time = wake_time.AddDate(0, 0, 1)
        }

        wake_time = time.Date(wake_time.Year(), wake_time.Month(),
            wake_time.Day(), hours, minutes, 0, 0, wake_time.Location())
    }

    first_time := wake_time.Add(-540 * time.Minute)
    second_time := wake_time.Add(-450 * time.Minute)
    third_time := wake_time.Add(-360 * time.Minute)
    fourth_time := wake_time.Add(-270 * time.Minute)


    times_fmt := "3:04 PM"
    fmt.Printf("To wake up at %v:\n", wake_time.Format(times_fmt))
    fmt.Printf("You should sleep at: %v\n\n", first_time.Format(times_fmt))
    fmt.Printf("Or yet, you could sleep at: %v\n", second_time.Format(times_fmt))
    fmt.Printf("Also at: %v\n", third_time.Format(times_fmt))
    fmt.Printf("Or at: %v\n", fourth_time.Format(times_fmt))
}
