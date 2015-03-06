package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Repo struct {
	URL string
	Path string
	Branch string

}

type Signature struct {
	Name string
	Email string
}

type Config struct {
	Repos map[string]Repo
	Signature
}

// Read YAML config file
func ReadConfig(filename string) (Config, error) {
	var config Config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

