package read_query

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"influx2/config"
	"influx2/internal/app"
)

var ctx = context.Background()

type readAgent struct {
	cfg config.InfluxDB
}

func (ra *readAgent) Run(client influxdb2.Client) error {
	fmt.Println("Welcome Read Query From Stdin Agent ...")

	reader := bufio.NewReader(os.Stdin)
	var queryBytes []byte

	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if string(buf) == "" {
			break
		}

		queryBytes = append(queryBytes, buf...)
		queryBytes = append(queryBytes, []byte("\n")...)
	}

	query := string(queryBytes)

	fmt.Printf("query:\n%s\n", query)

	result, err := client.QueryAPI(ra.cfg.OrgName).Query(ctx, query)
	if err != nil {
		fmt.Printf("query failed .err=%v\n", err)
		return err
	}

	//	view.ViewDetail(result)

	if result.Err() != nil {
		fmt.Printf("query parsing error: %v\n", result.Err())
		return err
	}
	return nil
}

func NewAgentService(cfg config.InfluxDB) app.AgentService {
	return &readAgent{cfg: cfg}
}
