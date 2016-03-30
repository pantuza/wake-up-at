

package main


import (
    "fmt"
    "time"
    "flag"
)


const MORNING string = "AM"


const MIDDAY int = 12


// Parses the command line options
func parseOptions(hours, minutes *int, period *string) {

    hours = flag.Int("h", 8, "Hour to wake up")
    minutes = flag.Int("m", 0, "Minute to wake up in the given hour")
    period = flag.String("p", "am", "Period of the day: am/pm")
    fmt.Printf("Hours", *hours)
    flag.Parse()
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

    diff := wake_time.Add(-270 * time.Minute)
    fmt.Printf("Wake time: %v\n", diff)
}
