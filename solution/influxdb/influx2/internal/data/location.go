package data

import (
	"fmt"
	"math/rand"

	"github.com/go-faker/faker/v4"
)

var (
	NodeNumbers      = []int32{0, 1, 2, 3, 4}
	Namespaces       = []string{"foo", "bar"}
	NamespaceNumbers = []int32{101, 102}
	AppNames         = []string{"alert", "metric", "proxy", "config", "api-server", "trace"}
	PodNumbers       = []int32{0, 1, 2}
	ContainerNames   = []string{"istio", "appName"}
)

type Location struct {
	NodeName      string
	Namespace     string
	AppName       string
	PodName       string
	ContainerName string
}

type LocationFaker struct {
	NodeNumber      int32  `faker:"oneof: 1, 2, 3, 4, 5"`
	NamespacePrefix string `faker:"oneof: foo, bar"`
	NamespaceNumber int32  `faker:"oneof: 101, 102"`
	AppName         string `faker:"oneof: alert, metric, proxy, config, api-server, trace"`
	PodNumber       int32  `faker:"oneof: 0,1,2"`
	ContainerName   string `faker:"oneof: istio, appName"`
}

func WokerNodeName() string {
	number := NodeNumbers[rand.Intn(len(NodeNumbers))]
	return fmt.Sprintf("worker-%d", number)
}

func NewLocation() []Location {
	var items []Location
	for _, ns := range Namespaces {
		for _, nsNumber := range NamespaceNumbers {
			realNs := fmt.Sprintf("%s-%d", ns, nsNumber)
			for _, appName := range AppNames {
				for _, podNumber := range PodNumbers {
					podName := fmt.Sprintf("%s-%d", appName, podNumber)
					nodeName := WokerNodeName()
					for _, containerName := range ContainerNames {
						l := Location{
							NodeName:  nodeName,
							Namespace: realNs,
							AppName:   appName,
							PodName:   podName,
						}
						if containerName == "appName" {
							l.ContainerName = appName
						} else {
							l.ContainerName = containerName
						}

						items = append(items, l)

					}
				}
			}
		}
	}
	return items
}

func NewLocationByFaker() Location {
	f := LocationFaker{}
	faker.FakeData(&f)
	l := Location{}
	l.NodeName = fmt.Sprintf("worker-%d", f.NodeNumber)
	l.Namespace = fmt.Sprintf("%s-%d\n", f.NamespacePrefix, f.NamespaceNumber)
	l.AppName = f.AppName
	l.PodName = fmt.Sprintf("%s-%d", l.AppName, f.PodNumber)
	if f.ContainerName == "appName" {
		l.ContainerName = f.AppName
	} else {
		l.ContainerName = f.ContainerName
	}

	return l
}

func NewSliceLocationByFaker(num int) []Location {
	var items []Location

	for i := 0; i < num; i++ {
		items = append(items, NewLocationByFaker())
	}
	return items
}
