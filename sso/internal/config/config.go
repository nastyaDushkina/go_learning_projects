package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"` //env-required
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// с must начинаются те ф-ции, которые не обрабатывают ошибку
// они честно паникуют
func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// получает путь до конфига либо из перемнной окр-я, либо из флага командной строки. Флаг в приоритете
func fetchConfigPath() string {
	var res string

	// --config="path/config.yaml"
	// CONFIG_PATH=./path/file.yaml sso
	flag.StringVar(&res, "config", "", "path to conffig file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
