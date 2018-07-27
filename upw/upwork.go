package upw

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/rustyeddy/logrus"

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
		getAccessToken(cfg)
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

func getAccessToken(cfg *api.Config) {
	aurl := client.GetAuthorizationUrl("")

	// read verifier
	reader := bufio.NewReader(os.Stdin)
	log.Debugln("Visit the authorization url and provide oauth_verifier for ")
	log.Debugln("further authorization")
	log.Debugln(aurl)
	verifier, _ := reader.ReadString('\n')

	// get access token
	token := client.GetAccessToken(verifier)
	log.Debug("authorization token", token)

	cfg.AccessToken = token.Token
	cfg.AccessSecret = token.Secret
	err := saveConfig(cfgfile, cfg)
	if err != nil {
		log.Fatalf("failed to save config file %v", err)
	}
}
