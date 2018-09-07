package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	configMu sync.RWMutex
	config   = map[string]string{}
)

func Setup(c map[string]string) error {
	for key, value := range c {
		Set(key, value)
	}

	return nil
}

// Get an value by key
func Get(key string) string {
	configMu.RLock()
	defer configMu.RUnlock()

	return config[key]
}

// Set an value by key
func Set(key, value string) {
	configMu.Lock()
	defer configMu.Unlock()

	config[key] = value
}

// Parse a config file
func Parse(path string) {
	data, err := ioutil.ReadFile(path)
	c := make(map[string]string)

	if err != nil {
		return
	}

	yaml.Unmarshal(data, &c)

	for key, value := range c {
		Set(key, value)
	}
}
