package utils

import (
    "fmt"
    "time"
)

var start time.Time

func ActualizeTimer() {
    start = time.Now()
}

func PrintDuration(msg string) {
    elapsed := time.Since(start)
    fmt.Printf(msg+" %s\n", elapsed)
}
