

package main


import (
    "fmt"
    "time"
)


const MORNING string = "AM"


const MIDDAY int = 12


// Returns true if the period of the day is Morning. Otherwise, false
func isMorning (period string) bool { return period == MORNING }


func main () {

    now := time.Now()


    // Fake input data
    period := "AM"
    hours := 8
    minutes := 15


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
