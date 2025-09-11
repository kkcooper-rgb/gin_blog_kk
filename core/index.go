package core

import (
	"gin_blog_kk/config"
	"gin_blog_kk/utils"
	"log"

	"gopkg.in/yaml.v3"
)

func InitConfig() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML configuration: %v", err)
	}
	return c
}
