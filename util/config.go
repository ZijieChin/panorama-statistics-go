package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

func GetConfig() (config Config) {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		config.Port = "8080" // Default port
		return
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		config.Port = "8080"
		return
	}
	return
}
