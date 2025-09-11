package utils

import "os"

const configFile = "config.ymal"

func LoadYAML() ([]byte, error) {
	return os.ReadFile(configFile)
}
