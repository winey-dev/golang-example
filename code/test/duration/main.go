package main

import (
	"flag"
	"fmt"
	"time"
)

// REALTIME 단위일 경우 10초 기준
// 1MIN 단위 일 경우 1분 단위
// 5MIN 단위 일 경우 5분 단위
func main() {
	// 입력받은 값 확인
	var s_interval, s_range string
	flag.StringVar(&s_interval, "i", "10s", "interval")
	flag.StringVar(&s_interval, "interval", "10s", "interval")
	flag.StringVar(&s_range, "r", "5m", "time range")
	flag.StringVar(&s_range, "range", "5m", "time range")
	flag.Parse()

	fmt.Println(s_interval, s_range)

	interval, err := time.ParseDuration(s_interval)
	if err != nil {
		fmt.Printf("%s interval parser failed. err=%v\n", s_interval, err)
		return
	}

	timeRange, err := time.ParseDuration(s_range)
	if err != nil {
		fmt.Printf("%s range parser failed. err=%v\n", s_range, err)
		return
	}

	// 구해야 하는 시간 값 갯수
	countTime := int64(timeRange) / int64(interval)
	fmt.Println(countTime)

	// 현재 시간 구하기
	curTime := time.Now()

	// 현재 시간에서 뒤에 한자리 숫자 초시간 없애기
	unixCurTime := curTime.UnixNano()
	unixCurTime = unixCurTime - (unixCurTime % int64(interval))

	times := make([]string, countTime)

	for i := countTime - 1; i >= 0; i-- {
		rt := time.Unix(0, unixCurTime)
		times[i] = rt.Format("15:04:05")
		unixCurTime = unixCurTime - int64(interval)
	}

	for i, t := range times {
		fmt.Println(i, t)
	}
}
