package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}

// env-default :"production"
type Config struct {
	ENV         string `yaml:"env" env:"ENV" env-require:"true" `
	StoragePath string `yaml:"storage_path"  env-require:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path tothe configuration file")
		flag.Parse()

		configPath = *flags

		if configPath ==""{
			log.Fatal("config pah is not set")
		}
	}

	if _, err := os.Stat(configPath) ; os.IsNotExist(err){
		log.Fatalf("config file does not exist %s" , configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath , &cfg)

	if err != nil{
		log.Fatalf("con not read config file: %s", err.Error())
	}
	return &cfg
}
