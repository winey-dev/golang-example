package main

import (
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05.000000"

func main() {
	now := time.Now()
	duration, _ := time.ParseDuration("15m")

	fmt.Printf("%s\n", now.Format(TimeFormat))
	fmt.Printf("%s\n", now.Add(-1*duration).Format(TimeFormat))

	duration = now.AddDate(0, 1, 0).Sub(now)
	fmt.Printf("%s\n", now.Add(-1*duration).Format(TimeFormat))

}
