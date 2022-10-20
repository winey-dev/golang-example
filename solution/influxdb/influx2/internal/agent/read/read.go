package read

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"influx2/config"
	"influx2/internal/app"
	"influx2/internal/pkg/fluxql"
)

var ctx = context.Background()

type readAgent struct {
	cfg config.InfluxDB
}

func (ra *readAgent) Run(client influxdb2.Client) error {
	fmt.Println("Welcome Read Default Agent ...")

	// InfluxDB 2.0  Query,
	query := fluxql.NewBuilder().
		From(ra.cfg.Bucket).
		RangeString("-1h").
		Measurement("sample-1").
		Field("CPU_USAGE").
		Builder().
		String()
	fmt.Printf("%s\n", query)

	result, err := client.QueryAPI(ra.cfg.OrgName).Query(ctx, query)
	if err != nil {
		fmt.Printf("query failed .err=%v\n", err)
		return err
	}

	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table position: %d\n", result.TableMetadata().Position())
		}

		values := result.Record().Values()
		for k, v := range values {
			if k == "table" {
				continue
			}
			fmt.Printf("{%s : %v} ", k, v)
		}
		fmt.Println()

	}

	if result.Err() != nil {
		fmt.Printf("query parsing error: %v\n", result.Err())
		return err
	}
	return nil
}

func NewAgentService(cfg config.InfluxDB) app.AgentService {
	return &readAgent{cfg: cfg}
}
