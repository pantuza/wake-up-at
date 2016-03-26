

package main


import (
    "fmt"
    "time"
)


func main () {

    now := time.Now()

    fmt.Printf("Hello nurse: %d\n", now.Year())
}
