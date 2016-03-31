
package main


import (
    "testing"
    "os"
    "flag"
)


// Reset CommandLine for each test from flags parsing
func ResetCommandLineFlags() {

    flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}


func TestDefaultCommandlineOptionsValues(t *testing.T) {

    defer ResetCommandLineFlags()

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

    defer ResetCommandLineFlags()

    os.Args = []string{"cmd", "-h", "10"}

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if hours != 10 {
        t.Fatal("wake up hour must be 10, got ", hours)
    }
}


func TestMinuteInputOption(t *testing.T) {

    defer ResetCommandLineFlags()

    os.Args = []string{"cmd", "-m", "15"}

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if minutes != 15 {
        t.Fatal("wake up minutes must be 15, got ", minutes)
    }
}


func TestPeriodInputOption(t *testing.T) {

    defer ResetCommandLineFlags()

    os.Args = []string{"cmd", "-p", "pm"}

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if period != "PM" {
        t.Fatal("wake up period must be PM, got ", period)
    }
}


func TestWrongPeriodInputOption(t *testing.T) {

    defer ResetCommandLineFlags()

    os.Args = []string{"cmd", "-p", "wrong"}

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if period != "AM" {
        t.Fatal("wake up period must be PM, got ", period)
    }
}


func TestUsageFunctionCallWhenThereIsAnInvalidOption(t *testing.T) {

    defer ResetCommandLineFlags()

    os.Args = []string{"cmd", "-wrong"}

    var hours int
    var minutes int
    var period string

    var called bool = false
    flag.Usage = func () { called = true }

    parseOptions(&hours, &minutes, &period)

    if !called {
        t.Fatal("Usage function not called")
    }
}
