
package main


import (

    "testing"
    "time"
)


func TestCalcTimes(t *testing.T) {

    base_time := time.Date(2016, time.April, 10, 8, 0, 0, 0, time.Local)

    var first_time, second_time, third_time, fourth_time time.Time
    calcTimes(&base_time, &first_time, &second_time, &third_time, &fourth_time)

    if first_time.Hour() != 23 {
        t.Fatal("First time shoulb be 23, got %d", first_time.Hour())
    } else if second_time.Hour() != 00 {
        t.Fatal("Second time shoulb be 00, got %d", second_time.Hour())
    } else if third_time.Hour() != 2 {
        t.Fatal("Third time shoulb be 2, got %d", third_time.Hour())
    } else if fourth_time.Hour() != 3 {
        t.Fatal("Fourth time shoulb be 3, got %d", fourth_time.Hour())
    }
}
