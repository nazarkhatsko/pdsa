package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Runners []RunnerConfig `yaml:"runners"`
	Viewers []ViewerConfig `yaml:"viewers"`
}

type RunnerConfig struct {
	Tag        string   `yaml:"tag"`
	Iterations int      `yaml:"iterations"`
	Solver     string   `yaml:"solver"`
	Strategies []string `yaml:"strategies"`
}

type ViewerConfig struct {
	Type string `yaml:"type"`
	Out  string `yaml:"out"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("error: config path is empty")
	}

	return MustLoadPath(path)
}

func MustLoadPath(path string) *Config {
	var config Config

	file := readConfigFile(path)
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		panic("error occurred while parsing config file: " + path)
	}

	return &config
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}

func readConfigFile(path string) []byte {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("error: config file with name \"%s\" does not exist", path))
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic("an error occurred while reading a config file: " + path)
	}

	return file
}
