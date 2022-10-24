package stat_map

import (
	"fmt"
	"sync"
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
	// ItemWithTag Key ItemName:map[k1:v1 k2:v2]
	ItemWithTag map[string]StatWithTag

	mutex *sync.Mutex
}

/*
response_faile  error_code: 500
response_faile  error_code: 400
*/

type Stat struct {
	ItemName  string
	ItemValue float64
}

type StatWithTag struct {
	ItemName  string
	ItemValue float64
	Tags      map[string]string
}

func NewStatInfo() *StatInfo {
	return &StatInfo{
		Item:        make(map[string]Stat),
		ItemWithTag: make(map[string]StatWithTag),
		mutex:       &sync.Mutex{},
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

func (s *StatInfo) IncreaseValue(itemName string, sum int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	stat, ok := s.Item[itemName]
	if !ok {
		s.Item[itemName] = Stat{
			ItemName:  itemName,
			ItemValue: float64(sum),
		}
	} else {
		temp := int(stat.ItemValue) + sum
		stat.ItemValue = float64(temp)
		s.Item[itemName] = stat
	}
}

func (s *StatInfo) Increase(itemName string) {
	s.IncreaseValue(itemName, 1)
}

func (s *StatInfo) DecreaseValue(itemName string, decrease int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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

func (s *StatInfo) Decrease(itemName string) {
	s.DecreaseValue(itemName, 1)
}

func (s *StatInfo) SetValue(itemName string, itemValue float64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Item[itemName] = Stat{
		ItemName:  itemName,
		ItemValue: itemValue,
	}
}

func (s *StatInfo) SetValueWithTag(itemName string, itemValue float64, tags map[string]string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.ItemWithTag[makeKey(itemName, tags)] = StatWithTag{
		ItemName:  itemName,
		ItemValue: itemValue,
		Tags:      tags,
	}
}

func makeKey(itemName string, tags map[string]string) string {
	return fmt.Sprintf("%s:%v", itemName, tags)
}

func (s *StatInfo) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
