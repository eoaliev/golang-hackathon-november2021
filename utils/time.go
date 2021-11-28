package utils

import (
    "fmt"
    "time"
)

var timers map[string]time.Time


func ActualizeTimer(name string) {
    if (timers == nil) {
        timers = make(map[string]time.Time)
    }

    timers[name] = time.Now()
}

func PrintDuration(name string, msg string) {
    if _, ok := timers[name]; !ok {
        return
    }

    elapsed := time.Since(timers[name])
    fmt.Printf(msg+" %s\n", elapsed)
}

func StopTimer(name string) {
    if _, ok := timers[name]; ok {
        delete(timers, name)
    }
}
