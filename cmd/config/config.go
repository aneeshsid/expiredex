package config

import (
	"fmt"
	"os"

	aero "github.com/aerospike/aerospike-client-go/v8"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Aerospike AerospikeConfig  `yaml:"aerospike"`
	Cleanup   AerospikeCleanUp `yaml:"cleanup"`
}

type AerospikeCleanUp struct {
	Key_Prefix  string `yaml:"key_prefix"`
	Date_Format string `yaml:"date_format"`
}

type AerospikeConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Set       string `yaml:"set"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
}

func ReadAerospikeConfig(filename string) (*AerospikeConfig, *AerospikeCleanUp, error) {
	configContent, err := os.ReadFile(filename)

	if err != nil {
		return nil, nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(configContent, &cfg)

	if err != nil {
		return nil, nil, err
	}
	return &cfg.Aerospike, &cfg.Cleanup, nil
}

func panicOnError(err error) (verdict bool) {
	return err != nil
}

func ClientConnect(config *AerospikeConfig) (client *aero.Client) {
	client, err := aero.NewClient(config.Host, config.Port)

	if panicOnError(err) {
		panic(err)
	}

	fmt.Println("Succesfully connected to aerospike client with hostname", config.Host, "port: ", config.Port)
	return client
}
