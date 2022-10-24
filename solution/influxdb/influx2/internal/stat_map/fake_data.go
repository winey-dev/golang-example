package stat_map

import (
	"encoding/json"
	"fmt"
	"influx2/internal/data"
)

func AllMakeStatInfo() []*StatInfo {
	var items []*StatInfo
	locs := data.NewLocation()
	for _, loc := range locs {
		s := NewStatInfo()
		s.SetLocation(loc.NodeName, loc.Namespace, loc.AppName, loc.PodName, loc.ContainerName)
		items = append(items, s)
	}
	return items
}

func (s *StatInfo) FakeData() {
	rscs := data.NewResourceFieldValue()
	for _, rsc := range rscs {
		s.SetValue(rsc.ItemName, rsc.UpdateValue())
	}

	reqresps := data.NewReqRespValue()

	if len(reqresps) > 1 {
		if reqresps[0].ItemValue < reqresps[1].ItemValue {
			fmt.Printf("new req resp value error data make\n")
			data, _ := json.Marshal(reqresps)
			fmt.Println(string(data))
			return
		}
	}

	for i, reqresp := range reqresps {
		if i < 2 {
			s.SetValue(reqresp.ItemName, float64(reqresp.ItemValue))
		} else {
			s.SetValueWithTag(reqresp.ItemName, float64(reqresp.ItemValue), map[string]string{"error_code": reqresp.Code})
		}
	}
}
