package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/upwork/golang-upwork/api"
)

// GetConfig for upwork client
func GetConfig() *api.ApiClient {
	cfg := api.ReadConfig(cfgFile)
	client := api.Setup(cfg)
	if !client.HasAccessToken() {
		getAccessToken()
	}
	return &client
}

// SaveConfig will write config to a File
func saveConfig(fn string, cfg *api.Config) error {

	// read from the config file
	d, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal json %v", err)
	}
	err = ioutil.WriteFile(fn, d, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file %s -> %v", fn, err)
	}
	return nil
}
