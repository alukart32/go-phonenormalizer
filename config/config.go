package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Port     int16  `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBname   string `json:"dbname" yaml:"dbname"`
}

type PhoneNormalizer struct {
	Timeout int64  `json:"timeout" yaml:"timeout"`
	Regex   string `json:"regex" yaml:"regex"`
}

type Config struct {
	DB              DBConfig        `json:"db" yaml:"db"`
	PhoneNormalizer PhoneNormalizer `json:"normalizer" yaml:"normalizer"`
}

// New returns the new config by filepath. A successful call returns err == nil.
// Errors can be caused by incorrect description of the file path or data structure.
func New(cfgPath string) (*Config, error) {
	log.Println("read the config file...")
	return readConfig(cfgPath)
}

// Print outputs the configuration in YAML format.
func Print(c *Config) {
	if data, err := yaml.Marshal(*c); err != nil {
		log.Println("can not print config")
	} else {
		log.Printf("config data\n%s%v", "---\n", string(data))
	}
}

func readConfig(path string) (*Config, error) {
	fail := func(err error) (*Config, error) {
		return nil, fmt.Errorf("Config: %v", err)
	}

	if len(path) == 0 {
		return fail(fmt.Errorf("path is empty"))
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".yaml", ".json":
	default:
		return fail(fmt.Errorf("unexpected extension"))
	}

	f, err := os.Open(path)
	if err != nil {
		return fail(err)
	}
	defer f.Close()

	cfg := &Config{}
	switch ext {
	case ".yaml":
		if err = yaml.NewDecoder(f).Decode(cfg); err != nil {
			return fail(err)
		}
	case ".json":
		if err = json.NewDecoder(f).Decode(cfg); err != nil {
			return fail(err)
		}
	}

	return cfg, nil
}
