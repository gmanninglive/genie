package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func LoadConfig(config_location string) Config {
  var config []Task
  f, err := os.ReadFile(config_location)
  Check(err)
  
  json.Unmarshal([]byte(f), &config)
  
  
  
  out := make(Config, len(config))

  for i := 0; i < len(config); i++ {
    config[i].Base = filepath.Dir(config_location)
    out[i] = config[i]
  }
  
  return out
}
