package main

import (
	"fmt"
	"time"
)

func main() {
	var Times []string
	var Times2 []string
	interval := time.Second * 10
	timeRange, _ := time.ParseDuration("5m")
	fmt.Println(int64(interval), int64(timeRange))
	fmt.Println(int64(timeRange) / int64(interval))

	t := time.Now().Unix()
	t = t - (t % 10)
	for i := 0; i < int(10); i++ {
		rt := time.Unix(t, 0)
		Times = append(Times, rt.Format("2006-01-02 15:04:05"))
		Times2 = append(Times2, rt.Format("15:04:05"))
		t = t - 10
	}

	for i, t := range Times {
		fmt.Println(i, t, Times2[i])
	}
}
