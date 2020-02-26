package config

import (
	"gopkg.in/yaml.v2"
	"github.com/kelseyhightower/envconfig"
	"os"
)

// Define the layout of the configuration file
type Config struct {
	Development struct {
		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"post"`
		} `yaml:"server"`
		MongoDB struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"mongodb"`
		Postgres struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"postgres"`
	} `yaml: "dev"`
	Staging struct {
		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"post"`
		} `yaml:"server"`
		MongoDB struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"mongodb"`
		Postgres struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"postgres"`
	} `yaml: "test"`
	Production struct {
		Server struct {
			Host string `yaml:"host"`
			Port string `yaml:"post"`
		} `yaml:"server"`
		MongoDB struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"mongodb"`
		Postgres struct {
			Username string `yaml:"user"`
			Password string `yaml:"pass"`
		} `yaml:"postgres"`
	} `yaml: "prod"`
}

func NewConfig() *Config {
	return &Config{}
}

// Parses existing config.yaml file for configuration
// settings.
func (c *Config) ParseConfigfile(filename string) {
	// Set variable to keep configuration v
	// Parse configuration into Cfg
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}
}

// Parses environment variables to overwrite the adhoc definition
// given in the config.yaml file
func (c *Config) ParseEnvVars() {
	err := envconfig.Process("", c)
	if err != nil {
		panic(err)
	}
}
