package write_field_item

import (
	"fmt"
	"influx2/config"
	"influx2/internal/app"
	"influx2/internal/data"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type writeAgent struct {
	cfg config.InfluxDB
}

func (wa *writeAgent) Run(client influxdb2.Client) error {
	fmt.Println("Welcome Write Field Item ...")
	WriteAPI := client.WriteAPI(wa.cfg.OrgName, wa.cfg.Bucket)

	errorsChan := WriteAPI.Errors()
	go func() {
		for err := range errorsChan {
			fmt.Printf("write field item error: %v\n", err)
		}
	}()

	fieldValues := data.NewResourceFieldValue()
	locations := data.NewLocation()

	timer := time.NewTicker(time.Second * 10)

	i := 0
	for {
		select {
		case <-timer.C:
			startTime := time.Now()
			for _, l := range locations {
				tags := map[string]string{
					"node_name":      l.NodeName,
					"namespace":      l.Namespace,
					"app_name":       l.AppName,
					"pod_name":       l.PodName,
					"container_name": l.ContainerName,
				}

				fields := map[string]interface{}{}
				for _, fv := range fieldValues {
					fields[fv.ItemName] = fv.UpdateValue()
				}

				p := write.NewPoint(
					"resource",
					tags,
					fields,
					time.Now(),
				)

				WriteAPI.WritePoint(p)
			}

			WriteAPI.Flush()
			// Flush
			endTime := time.Since(startTime)
			elapsedTime := float64(endTime.Milliseconds()) / float64(1000)
			fmt.Printf("[%d] insert position count:%d, elapsed:%0.3fms\n", i, len(locations)*len(fieldValues), elapsedTime)
		}
		i++
	}
}

func NewAgentService(cfg config.InfluxDB) app.AgentService {
	return &writeAgent{cfg: cfg}
}
