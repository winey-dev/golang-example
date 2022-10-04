package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	Server struct {
		LogLevel string `mapstructure:"logLevel"`
		Address  string `mapstructure:"address"`
		Port     string `mapstructure:"port"`
	}

	Config struct {
		v      *viper.Viper `yaml:"-"`
		env    string
		Server Server `yaml:"server"`
	}
)

func LoadConfig(env string) *Config {
	c := Config{
		env: env,
		v:   viper.New(),
	}

	c.v.SetConfigName("conf")
	c.v.AddConfigPath("./config/")
	c.v.SetConfigType("yaml")
	c.v.AutomaticEnv()

	if err := c.v.ReadInConfig(); err != nil {
		fmt.Println("read config failed. ", err)
		os.Exit(1)
	}

	c.loadConfig()

	c.v.WatchConfig()
	c.v.OnConfigChange(func(e fsnotify.Event) {
		c.reloadConfig()
	})
	return &c
}

func (c *Config) loadConfig() {
	if err := c.v.Unmarshal(c); err != nil {
		fmt.Println("viper read in config next step <unmarshal failed. err=%v\n", err)
		os.Exit(1)
	}
}

func (c *Config) reloadConfig() {
	// not exit config
	if c.Server.LogLevel != c.v.GetString("server.logLevel") {
		fmt.Printf("change server.logLevel %s->%s\n", c.Server.LogLevel, c.v.GetString("server.logLevel"))
		c.Server.LogLevel = c.v.GetString("server.logLevel")
	}

	// exit config
	if c.Server.Address != c.v.GetString("server.address") {
		fmt.Printf("change server.address %s->%s\n", c.Server.Address, c.v.GetString("server.address"))
		c.rutine("server.address")
		c.Server.Address = c.v.GetString("server.address")
	}

	// exit config
	if c.Server.Port != c.v.GetString("server.port") {
		fmt.Printf("change server.port %s->%s\n", c.Server.Port, c.v.GetString("server.port"))
		c.rutine("server.port")
		c.Server.Port = c.v.GetString("server.port")
	}
}

func (c *Config) rutine(configKey string) {
	if c.env == "local" {
		fmt.Printf("if you want to apply %s, you need to reboot\n", configKey)
	} else {
		fmt.Printf("reboot for the %s to appy\n", configKey)
		os.Exit(0)
	}
}

func (c *Config) Print() {
	fmt.Printf("LOG LEVEL      : %s\n", c.Server.LogLevel)
	fmt.Printf("SERVER ADDRESS : %s\n", c.Server.Address)
	fmt.Printf("SERVER PORT    : %s\n", c.Server.Port)
}
