package config

import (
	"github.com/parking_lot/src/constants"
	"os"
	"encoding/json"
	"log"
)

var (
	cfg *Config
)

func Init() {
	cfg = &Config{}
	err := readConfig(cfg, constants.ConfigFilePath,constants.Module)
	if err != nil {
		log.Fatalln("fatal : [Config][Init] failed to read config", err)
	}
}

func Get() *Config {
	if cfg == nil {
		Init()
	}
	return cfg
}

func readConfig(cfg *Config, path string, module string) error{
	environ := os.Getenv("PARKINGLOT")
	if environ == "" {
		environ = "development"
	}

	filename := path + "/" + module + "." + environ + ".json"
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}
	return nil
}