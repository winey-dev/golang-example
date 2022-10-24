package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		v        *viper.Viper `yaml:"-"`
		InfluxDB InfluxDB     `yaml:"influxdb"`
	}
	InfluxDB struct {
		EndPoint string `yaml:"endpoint"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		OrgName  string `yaml:"orgname"`
		OrgID    string `yaml:"orgid"`
		Token    string `yaml:"token"`
		Bucket   string `yaml:"bucket"`
	}
)

func LoadConfig() *Config {
	cfg := new(Config)
	cfg.v = viper.New()

	cfg.v.SetConfigName(fmt.Sprintf("setting"))
	cfg.v.AddConfigPath("../config")
	cfg.v.SetConfigType("yaml")

	if err := cfg.v.ReadInConfig(); err != nil {
		fmt.Printf("read config failed. err=%v\n", err)
		os.Exit(1)
	}

	if err := cfg.v.Unmarshal(cfg); err != nil {
		fmt.Printf("unmarshal failed. err=%v\n", err)
		os.Exit(1)
	}

	return cfg
}
