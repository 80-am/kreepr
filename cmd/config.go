package cmd

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config to use
type Config struct {
	DbUser		 string `yaml:"user"`
	DbPassword	 string `yaml:"password"`
	DbSchema	 string `yaml:"schema"`
	Key          string `yaml:"key"`
	Secret       string `yaml:"secret"`
	Token        string `yaml:"token"`
	AccessToken  string `yaml:"access_token"`
	AccessSecret string `yaml:"access_secret"`
}

// GetConfig of user
func (c *Config) GetConfig(secrets string) *Config {
	if secrets == "" {
		secrets = "secrets.yml"
	}
	yamlFile, err := ioutil.ReadFile(secrets)
	if err != nil {
		log.Printf("Error in %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}