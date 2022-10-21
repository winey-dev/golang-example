package stat_map

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func TestAllMakeStat(t *testing.T) {

	var points []*write.Point

	all := AllMakeStatInfo()
	for _, item := range all {
		item.FakeData()
		iPoints := item.NewPoint()
		fmt.Printf("point:%d\n", len(iPoints))
		points = append(points, iPoints...)

	}

	fmt.Printf("all point:%d\n", len(points))
}

func TestMakeStat(t *testing.T) {
	all := AllMakeStatInfo()
	statInfo := all[0]
	statInfo.FakeData()

	data, _ := json.Marshal(statInfo)

	fmt.Println(string(data))

}
