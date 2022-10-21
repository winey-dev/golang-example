package stat_array

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
	Item          []Stat
	ItemWithTag   []Stat
}

type Stat struct {
	RecordTime time.Time
	ItemName   string
	ItemValue  float64
	Tags       map[string]string
}

func NewStatInfo() *StatInfo {
	return new(StatInfo)
}

func (s *StatInfo) SetLocation(nodeName, namespace, appName, podName, containerName string) {
	s.NodeName = nodeName
	s.Namespace = namespace
	s.AppName = appName
	s.PodName = podName
	s.ContainerName = containerName
	return
}

func (s *StatInfo) SetValue(recordTime time.Time, itemName string, itemValue float64) {
	s.Item = append(s.Item, Stat{
		RecordTime: recordTime,
		ItemName:   itemName,
		ItemValue:  itemValue,
	})
}

func (s *StatInfo) SetValueWithTag(recordTime time.Time, itemName string, itemValue float64, tags map[string]string) {
	s.ItemWithTag = append(s.ItemWithTag, Stat{
		RecordTime: recordTime,
		ItemName:   itemName,
		ItemValue:  itemValue,
		Tags:       tags,
	})
}

func (s *StatInfo) Clear() {
	s.Item = nil
	s.ItemWithTag = nil

}

func (s *StatInfo) FakeData() {
	rscs := data.NewResourceFieldValue()
	for _, rsc := range rscs {
		s.SetValue(time.Now(), rsc.ItemName, rsc.UpdateValue())
	}
	reqresps := data.NewReqRespValue()
	for i, reqresp := range reqresps {
		if i < 2 {
			s.SetValue(time.Now(), reqresp.ItemName, float64(reqresp.ItemValue))
		} else {
			s.SetValueWithTag(time.Now(), reqresp.ItemName, float64(reqresp.ItemValue), map[string]string{"error_code": reqresp.Code})
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
	for _, item := range s.Item {
		fields := map[string]interface{}{
			item.ItemName: item.ItemValue,
		}

		point := write.NewPoint(
			s.Namespace,
			tags,
			fields,
			item.RecordTime,
		)
		points = append(points, point)
	}
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
			item.RecordTime,
		)
		points = append(points, point)
	}
	return points

}
