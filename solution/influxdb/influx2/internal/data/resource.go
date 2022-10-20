package data

import (
	"fmt"
	"strings"

	"github.com/go-faker/faker/v4"
)

// Make Resource Data
var (
	ResourceStats = []string{
		"cpu",
		"mem",
		"disk",
		"ni-rx",
		"ni-tx",
	}

	Units = []string{
		"usage",
		"use-by-bytes",
	}
)

type ResourceValue struct {
	ItemName  string
	FieldName string
	ItemValue float64
}

type ResourceValueFaker struct {
	ItemPrefix string `faker:"oneof:cpu,mem,disk,ni-tx,ni-rx"`
	ItemUnit   string `faker:"oneof:usage,use-by-bytes"`
}

type ValueFaker struct {
	Persent float64 `faker:"boundary_start=0.00, boundary_end=100.00"`
	Uint32  uint32
}

func NewResourceFieldValue() []ResourceValue {
	var items []ResourceValue

	for _, stat := range ResourceStats {
		for _, unit := range Units {
			items = append(items, ResourceValue{ItemName: fmt.Sprintf("%s-%s", stat, unit)})
		}
	}
	return items
}

func NewResourceTagValue() []ResourceValue {
	var items []ResourceValue

	for _, stat := range ResourceStats {
		for _, unit := range Units {
			items = append(items, ResourceValue{ItemName: stat, FieldName: unit})
		}
	}
	return items
}
func NewResourceValueByFaker() ResourceValue {
	f := ResourceValueFaker{}
	faker.FakeData(&f)
	return ResourceValue{
		ItemName: fmt.Sprintf("%s-%s", f.ItemPrefix, f.ItemUnit),
	}
}

func NewSliceResourceValueByFaker(num int) []ResourceValue {
	var items []ResourceValue

	for i := 0; i < num; i++ {
		items = append(items, NewResourceValueByFaker())
	}
	return items

}

func (f *ResourceValue) UpdateValue() float64 {
	vf := ValueFaker{}
	faker.FakeData(&vf)

	if f.FieldName == "" {
		if strings.Contains(f.ItemName, "usage") {
			f.ItemValue = vf.Persent
		} else {
			f.ItemValue = float64(vf.Uint32)
		}
	} else {
		if f.FieldName == "usage" {
			f.ItemValue = vf.Persent
		} else {
			f.ItemValue = float64(vf.Uint32)
		}
	}
	return f.ItemValue
}
