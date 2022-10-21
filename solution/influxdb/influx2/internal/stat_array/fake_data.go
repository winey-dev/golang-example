package stat_array

import "influx2/internal/data"

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
