package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 
func ParseFile(path string) (*DeviceConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var cfg DeviceConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML from %s: %w", path, err)
	}
	
	return &cfg, nil
}