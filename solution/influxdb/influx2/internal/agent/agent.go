package agent

import (
	"influx2/config"
	"influx2/internal/agent/read"
	"influx2/internal/agent/read_file"
	"influx2/internal/agent/read_query"
	"influx2/internal/agent/write_field_item"
	"influx2/internal/agent/write_tag_item"
	"influx2/internal/app"
	"influx2/internal/pkg/argument"
)

func NewAgentService(arg argument.Arguments, cfg config.InfluxDB) app.AgentService {
	if arg.Write {
		if arg.W_field {
			return write_field_item.NewAgentService(cfg)
		} else if arg.W_tag {
			return write_tag_item.NewAgentService(cfg)
		}
	} else {
		if arg.ReadFile != "" {
			return read_file.NewAgentService(cfg, arg.ReadFile)
		} else if arg.ReadQuery {
			return read_query.NewAgentService(cfg)
		}
	}

	return read.NewAgentService(cfg)
}
