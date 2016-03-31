
package main


import (
    "testing"
    "os"
    "flag"
)


func TestDefaultCommandlineOptionsValues(t *testing.T) {

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if hours != 8 {
        t.Fatal("Default hour must be 8, got ", hours)
    } else if minutes != 0 {
        t.Fatal("Default minute must be 0, got ", minutes)
    } else if period != "AM" {
        t.Fatal("Default period must be AM, got ", period)
    }
}


func TestHourInputOption(t *testing.T) {

    flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
    os.Args = []string{"cmd", "-m", "10"}

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if minutes != 10 {
        t.Fatal("wake up minutes must be 10, got ", minutes)
    }
}
