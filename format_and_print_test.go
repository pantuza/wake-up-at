
package main


import (

    "testing"
    "time"
)


func TestFormatAndPrint(t *testing.T) {

    base_time := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

    first_time := base_time.Add(-540 * time.Minute)
    second_time := base_time.Add(-450 * time.Minute)
    third_time := base_time.Add(-360 * time.Minute)
    fourth_time := base_time.Add(-270 * time.Minute)

    formatAndPrint(&base_time, &first_time, &second_time,
                   &third_time, &fourth_time)
    // Output:
    // To wake up at 8:00 AM, you should sleep at: 11:00 PM
    //
    // Also at: 12:30 AM | 2:00 AM | 3:30 AM
}
