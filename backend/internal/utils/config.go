package utils

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    ServerAddress string `yaml:"server_address"`
}

func LoadConfig(path string) (*Config, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}