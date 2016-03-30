
package main


import (
    "testing"
)


func TestDefaultCommandlineOptionsValues(t *testing.T) {

    var hours int
    var minutes int
    var period string

    parseOptions(&hours, &minutes, &period)

    if !(hours == 8) {
        t.Fatal("Default hour must be 8, got ", hours)
    }
}
