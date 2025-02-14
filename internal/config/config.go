package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local" env-required:"true"`
	// StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer          `yaml:"http_server"`
	PostgresDB          `yaml:"postgres_db"`
	RedisDB             `yaml:"redis_db"`
	GRPC_TInvest_server `yaml:"gRPC_TInvest_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type PostgresDB struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
	SSLmode  string `yaml:"sslmode"`
}

type RedisDB struct {
	Address     string        `yaml:"address"`
	Password    string        `yaml:"password"`
	Username    string        `yaml:"username"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

type GRPC_TInvest_server struct {
	Address  string `yaml:"address" env-default:"localhost:443"`
	SAddress string `yaml:"sandbox_address" env-default:"localhost:443"`
	Token    string `yaml:"token"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH_TInvest")
	// configPath := "E:\\project\\T_invest_api\\config\\local.yaml"
	if configPath == "" {
		log.Fatal("CONFIG_PATH_TInvest is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
