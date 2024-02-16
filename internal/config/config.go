package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	Storage  string        `yaml:"storage_path" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl: env-required:"true"`
	GRPC     GRPCConfig    `yaml:"grpc"`
}
type GRPCConfig struct {
	Port     int           `yaml:"port"`
	Timeoute time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {

	path := fetchConfgiPath()
	if path == "" {
		panic("empty path to conf")
	}
	if _, err := os.Stat(path); err != nil {
		panic("config file doesnt exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config")
	}
	return &cfg
}

func fetchConfgiPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to configfile")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
