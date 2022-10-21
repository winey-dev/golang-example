package stat_map

import (
	"influx2/internal/data"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type StatInfo struct {
	NodeName      string
	Namespace     string
	AppName       string
	PodName       string
	ContainerName string
	Item          map[string]Stat
	ItemWithTag   map[string]Stat
}

type Stat struct {
	ItemName  string
	ItemValue float64
	Tags      map[string]string
}

func NewStatInfo() *StatInfo {
	return &StatInfo{
		Item:        make(map[string]Stat),
		ItemWithTag: make(map[string]Stat),
	}
}

func (s *StatInfo) SetLocation(nodeName, namespace, appName, podName, containerName string) {
	s.NodeName = nodeName
	s.Namespace = namespace
	s.AppName = appName
	s.PodName = podName
	s.ContainerName = containerName
	return
}

func (s *StatInfo) SetValue(itemName string, itemValue float64) {
	s.Item[itemName] = Stat{
		ItemName:  itemName,
		ItemValue: itemValue,
	}
}

func (s *StatInfo) IncreaseValue(itemName string, increase int) {
	stat, ok := s.Item[itemName]
	if !ok {
		s.Item[itemName] = Stat{
			ItemName:  itemName,
			ItemValue: float64(increase),
		}
	} else {
		temp := int(stat.ItemValue) + increase
		stat.ItemValue = float64(temp)
		s.Item[itemName] = stat
	}
}

func (s *StatInfo) DecreaseValue(itemName string, decrease int) {
	stat, ok := s.Item[itemName]
	if !ok {
		s.Item[itemName] = Stat{
			ItemName:  itemName,
			ItemValue: float64(decrease),
		}
	} else {
		temp := int(stat.ItemValue) - decrease
		stat.ItemValue = float64(temp)
		s.Item[itemName] = stat
	}
}

func (s *StatInfo) SetValueWithTag(itemName string, itemValue float64, tags map[string]string) {
	s.ItemWithTag[itemName] = Stat{
		ItemName:  itemName,
		ItemValue: itemValue,
		Tags:      tags,
	}
}

func (s *StatInfo) Clear() {
	for k, _ := range s.Item {
		delete(s.Item, k)
	}

	for k, v := range s.ItemWithTag {
		for kk, _ := range v.Tags {
			delete(v.Tags, kk)
		}
		delete(s.ItemWithTag, k)
	}
}

func (s *StatInfo) FakeData() {
	rscs := data.NewResourceFieldValue()
	for _, rsc := range rscs {
		s.SetValue(rsc.ItemName, rsc.UpdateValue())
	}

	reqresps := data.NewReqRespValue()
	for i, reqresp := range reqresps {
		if i < 2 {
			s.SetValue(reqresp.ItemName, float64(reqresp.ItemValue))
		} else {
			s.SetValueWithTag(reqresp.ItemName, float64(reqresp.ItemValue), map[string]string{"error_code": reqresp.Code})
		}
	}
}

func (s *StatInfo) NewPoint() []*write.Point {
	var points []*write.Point
	points = append(points, s.NewPointWithItem()...)
	points = append(points, s.NewPointWithItemWithTag()...)
	return points
}

func (s *StatInfo) NewPointWithItem() []*write.Point {
	var points []*write.Point
	tags := map[string]string{
		"node_name":      s.NodeName,
		"app_name":       s.AppName,
		"pod_name":       s.PodName,
		"container_name": s.ContainerName,
	}
	fields := map[string]interface{}{}
	for _, item := range s.Item {
		fields[item.ItemName] = item.ItemValue
	}

	point := write.NewPoint(
		s.Namespace,
		tags,
		fields,
		time.Now(),
	)
	points = append(points, point)
	return points
}

func (s *StatInfo) NewPointWithItemWithTag() []*write.Point {
	var points []*write.Point

	for _, item := range s.ItemWithTag {
		tags := map[string]string{
			"node_name":      s.NodeName,
			"app_name":       s.AppName,
			"pod_name":       s.PodName,
			"container_name": s.ContainerName,
		}

		for k, v := range item.Tags {
			tags[k] = v
		}

		fields := map[string]interface{}{
			item.ItemName: item.ItemValue,
		}

		point := write.NewPoint(
			s.Namespace,
			tags,
			fields,
			time.Now(),
		)
		points = append(points, point)
	}

	return points

}
