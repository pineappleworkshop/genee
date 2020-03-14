package services

import (
	"io/ioutil"

	"genee/models"

	"gopkg.in/yaml.v2"
)

func ParseConfig(path string) (*models.Config, error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &models.Config{}
	if err := yaml.Unmarshal(configFile, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
