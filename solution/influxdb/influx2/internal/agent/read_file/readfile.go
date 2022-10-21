package read_file

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"influx2/config"
	"influx2/internal/app"
	"influx2/internal/pkg/table"
)

var ctx = context.Background()

type readAgent struct {
	cfg  config.InfluxDB
	path string
}

func (ra *readAgent) Run(client influxdb2.Client) error {
	fmt.Println("Welcome Read Query From File Agent ...")
	queryLine, err := ReadFromFile(ra.path)
	if err != nil {
		fmt.Printf("read file failed. err=%v\n", err)
		return err
	}

	for _, query := range queryLine {
		fmt.Printf("Run Query:\n%s\n", query)
		result, err := client.QueryAPI(ra.cfg.OrgName).Query(ctx, query)
		if err != nil {
			fmt.Printf("query failed .err=%v\n", err)
			continue
		}

		tables := table.Store(result)
		for _, t := range tables {
			t.Show()
		}

		if result.Err() != nil {
			fmt.Printf("query parsing error: %v\n", result.Err())
			continue
		}

	}

	return nil
}

func NewAgentService(cfg config.InfluxDB, path string) app.AgentService {
	return &readAgent{cfg: cfg, path: path}
}

func ReadFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("open failed. err=%v\n", err)
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var str string
	var queryLine []string
	var anno bool
	for {
		buf, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("readLine failed. err=%v\n", err)
			return nil, err
		}

		temp := strings.TrimLeft(string(buf), " ")
		if temp == "" {
			continue
		}

		r := []rune(temp)
		if r[0] == '#' {
			continue
		}

		if anno {
			if temp == "*/" {
				anno = false
			}
			continue
		}

		if temp == "/*" {
			anno = true
			continue
		}

		if strings.Contains(temp, "---") {
			if str != "" {
				queryLine = append(queryLine, str)
			}
			str = ""
			continue
		}

		str += string(buf) + "\n"
	}

	if str != "" {
		queryLine = append(queryLine, str)
	}

	return queryLine, nil

}
