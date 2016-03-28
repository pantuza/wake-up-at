

package main


import (
    "fmt"
    "time"
)


func main () {

    now := time.Now()

    // Fake input data
    in_period := "AM"
    in_hours := 8
    in_minutes := 15


    // Create the time to wake
    var hours time.Duration = time.Duration(in_hours) * time.Hour
    var minutes time.Duration = time.Duration(in_minutes) * time.Minute
    var wake_duration time.Duration = hours + minutes



    wake_time := time.Now()
    if in_period == "AM" {

        if now.Hour() > 12 {


            fmt.Printf("Duration time: %v\n", wake_duration)
            wake_time = wake_time.AddDate(0, 0, 1)
            //wake_time = wake_time.Truncate(wake_duration)
        }
    }

    fmt.Printf("Wake time: %v\n", wake_time)
    fmt.Printf("Hello nurse: %d\n", now.Hour())
}
