package upw

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/upwork/golang-upwork/api"
)

var (
	cfgfile string
	client  api.ApiClient
)

func init() {
	cfgfile = "/Users/rusty/.config/upwork.json"
}

// GetConfig for upwork client
func apiClient() api.ApiClient {
	cfg := api.ReadConfig(cfgfile)
	client := api.Setup(cfg)
	if !client.HasAccessToken() {
		// Get access token
		aurl := client.GetAuthorizationUrl("")
		fmt.Printf("  authorization URL %s", aurl)
		reader := bufio.NewReader(os.Stdin)
		verifier, _ := reader.ReadString('\n')
		token := client.GetAccessToken(verifier)
		fmt.Printf("   ~ Token: %v", token)
	}
	return client
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
