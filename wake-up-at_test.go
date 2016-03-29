package main


import (

    "testing"
)


func TestIsMorningTrue(t *testing.T) {

    if !isMorning("AM") {
        t.Fatal("Expected isMorning to true")
    }
}



func TestIsMorningFalse(t *testing.T) {

    if isMorning("PM") {
        t.Fatal("Expected isMorning to false")
    }
}



func TestIsMorningFalseWhenDurationIsInvalid(t *testing.T) {

    if isMorning("MM") {
        t.Fatal("Expected isMorning to false")
    }
}
