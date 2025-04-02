package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/charlesfan/task-go/utils/log"
)

type (
	EnvType  string
	LogLevel string
)

var content Config

const (
	Dev EnvType = "development"
	Pro EnvType = "production"

	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
)

func (e EnvType) String() string { return string(e) }

func (l LogLevel) String() string { return string(l) }

type Config struct {
	Env    string
	Store  string
	Server *Server
	Log    *Log
	Redis  *Redis
}

func New() Config {
	return content
}

func Init() {
	c := Config{
		Env:   "development",
		Store: "redis",
		Log: &Log{
			Level: DebugLevel.String(),
		},
		Server: &Server{
			Schema: "http",
			Host:   "0.0.0.0",
			Port:   "8080",
		},
		Redis: &Redis{
			Addr:     "store:6379",
			Password: "",
			DB:       0,
			PoolSize: 100,
		},
	}

	c.logInit()

	if err := c.loadConfig(); err != nil {
		log.Error("load config failed")
		os.Exit(1)
	}

	content = c
}

func (c *Config) loadConfig() error {
	cfgPath := viper.GetString("config")
	viper.SetEnvPrefix("TASKGO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Errorf("Error reading config file (%s), %v", cfgPath, err)
			os.Exit(1)
		}
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		log.Errorf("Unable to decode into struct, %v", err)
		return err
	}

	b, _ := json.MarshalIndent(&c, "", "  ")
	fmt.Println(string(b))

	return nil
}

func (c *Config) logInit() {
	log.Init(c.Env, c.Log.File, c.Log.Level)
}
