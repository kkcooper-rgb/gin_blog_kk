package core

import (
	"fmt"
	"gin_blog_kk/config"
	"gin_blog_kk/utils"

	"gopkg.in/yaml.v3"
)

func InitConfig() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		fmt.Println(err)
	}
	return c
}
