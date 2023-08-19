package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Ldap     LdapConfig     `yaml:"ldap"`
}

type ServerConfig struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LdapConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Dn   string `yaml:"dn"`
}

func GetConfig() (Config, error) {

	var config Config

	exePath, err := os.Executable()

	if err != nil {
		return config, err
	}

	exeDir := filepath.Dir(exePath)

	yamlFile, err := os.ReadFile(exeDir + "/config.yaml")

	if err != nil {
		return config, fmt.Errorf("error reading YAML file: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("error decoding YAML: %w", err)
	}

	return config, nil

}
