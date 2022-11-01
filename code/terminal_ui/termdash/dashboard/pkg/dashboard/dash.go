package dashboard

import (
	"fmt"
	"math/rand"
	"time"
)

type Dashboard struct {
	Name     string
	Charts   []*Chart
	Interval time.Duration
	Range    time.Duration
}

type Chart struct {
	Name   string
	Cols   []string    // X
	Values [][]float64 // X,Y
	Times  []string    // Y
}

var Data = Dashboard{
	Name:     "TEST",
	Interval: (time.Second * 10),
	Range:    (time.Second * 120), // 2m
	Charts: []*Chart{
		&Chart{
			Name: "NS-1-RESOURCE",
			Cols: []string{"CPU_USAGE", "MEM_USAGE", "DISK_USAGE", "NI_TX_USAGE", "NI_RX_USAGE"},
		},
		&Chart{
			Name: "NS-2-RESOURCE",
			Cols: []string{"CPU_USAGE", "MEM_USAGE", "DISK_USAGE", "NI_TX_USAGE", "NI_RX_USAGE"},
		},

		&Chart{
			Name: "NS-1-SUCC_RATE",
			Cols: []string{"REQUEST_SUCC_RATE", "ACCESS_SUCC_RATE", "QUERY_SUCC_RATE"},
		},
		&Chart{
			Name: "NS-2-SUCC_RATE",
			Cols: []string{"REQUEST_SUCC_RATE", "ACCESS_SUCC_RATE", "QUERY_SUCC_RATE"},
		},
	},
}

const timeFormat = "15:04:05"

func (c *Chart) InitValue(interval, timeRange time.Duration) {
	div := int64(timeRange) / int64(interval)
	c.Values = make([][]float64, len(c.Cols))
	for i := 0; i < len(c.Cols); i++ {
		c.Values[i] = make([]float64, div)
		for j := 0; j < int(div); j++ {
			c.Values[i][j] = float64(rand.Intn(100.00))
			// CPU : 0 1 2 3 4 5 6 7 // 각 시간대별로 들어감
		}
	}

	c.Times = make([]string, div)

	t := time.Now().UnixNano()
	t = t - (t % int64(interval))
	for i := div - 1; i >= 0; i-- {
		rt := time.Unix(0, t)
		c.Times[i] = rt.Format("15:04:05")
		t = t - int64(interval)
	}
}

func (c *Chart) RotateFloats() {
	// string to time
	// add interval
	// set time and value
	for i := 0; i < len(c.Values); i++ {
		c.Values[i] = c.Values[i][1:]
		c.Values[i] = append(c.Values[i], float64(rand.Intn(100.00)))
	}

	c.Times = c.Times[1:]
	newTime := time.Now().Unix()
	newTime = newTime - (newTime % 10)
	curTime := time.Unix(newTime, 0)
	c.Times = append(c.Times, curTime.Format(timeFormat))
}

func (c *Chart) MakeLabel() map[int]string {
	ret := make(map[int]string)
	for i, str := range c.Times {
		ret[i] = str
	}
	return ret
}

/*
type Table struct {
	Cols  []string    // X
	Time  []string    // Y
	Value [][]float64 // X,Y
}
*/

func (c *Chart) TableView() {
	// i : y
	// j : x

	fmt.Printf("%-10s  ", "Time")
	for i := 0; i < len(c.Cols); i++ {
		fmt.Printf("%-20s  ", c.Cols[i])
	}
	fmt.Println()
	fmt.Println("=============================================================================================================================================================")

	for y := 0; y < len(c.Times); y++ {
		fmt.Printf("%-10s  ", c.Times[y])
		for x := 0; x < len(c.Cols); x++ {
			fmt.Printf("%-20s  ", fmt.Sprintf("%0.3f", c.Values[x][y]))
		}
		fmt.Println()
	}
}
