package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	Filepath string `json:"-"`
	Hosts    []Host
}

func LoadConfig() (*Config, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(user.HomeDir, ".damn-butler")

	if ConfigExists(configPath) {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, err
		}

		buf := bytes.NewBuffer(data)
		var config Config
		err = json.NewDecoder(buf).Decode(&config)
		if err != nil {
			return nil, err
		}

		config.Filepath = configPath

		return &config, nil
	} else {
		return &Config{Filepath: configPath}, nil
	}
}

func ConfigExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func (c *Config) AddHost(host string) error {
	newHost, err := NewHost(host)
	if err != nil {
		return err
	}

	c.Hosts = append(c.Hosts, newHost)
	err = c.Save()
	if err != nil {
		fmt.Println("Failed to save config")
	}

	fmt.Println("Added", host, "to host list")

	return nil
}

func (c *Config) HasHosts() bool {
	return len(c.Hosts) > 0
}

func (c *Config) Save() error {
	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.Filepath, []byte(json), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) PrintHostList() {
	for _, h := range c.Hosts {
		fmt.Println(h.CCUrl)
	}
}
