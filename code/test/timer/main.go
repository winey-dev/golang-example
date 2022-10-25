package main

import (
	"fmt"
	"time"
)

/*
Handler 수행 시간 기준으로 Ticker가 발동
단, Handler 수행 시간이 Ticker 주기보다 길 경우, Ticker 호출이 밀리게 됨
*/

func Print(num int) {
	fmt.Printf("[%s]hello %d\n", time.Now().Format(time.RFC3339), num)
	time.Sleep(time.Second * 5)
	fmt.Printf("by %d\n", num)
}
func main() {
	timer := time.NewTicker(time.Second * 3)
	for {
		i := 0
		select {
		case <-timer.C:
			Print(i)
			i++
		}
	}
}
