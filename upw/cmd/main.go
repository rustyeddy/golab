package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rustyeddy/golab/upw"
	log "github.com/rustyeddy/logrus"

	"github.com/upwork/golang-upwork/api"
)

type request struct {
	query  string
	title  string
	skills string
}

var (
	cfgFile string = "/Users/rusty/.config/upwork.json"
	cfg     *api.Config
	client  api.ApiClient

	err error   // A catch all
	req request // request something from upwork
)

func init() {
	flag.StringVar(&req.query, "query", "", "General query to send upwork")
	flag.StringVar(&req.title, "title", "", "Search for pattern in title")
	flag.StringVar(&req.skills, "skills", "", "Search for skills")
}

func main() {
	var (
		jobs *upw.Jobs
		err  error
	)

	flag.Parse()
	client = getConfig()

	params := make(map[string]string)
	if req.query != "" {
		params["q"] = req.query
	}
	if req.title != "" {
		params["t"] = req.title
	}
	if req.skills != "" {
		params["s"] = req.skills
	}
	if len(params) < 1 {
		log.Fatal("at least one of query, title or skills are required")
	}

	if jobs, err = upw.FetchJobs(params); err != nil {
		log.Fatal("FetchJobs params: ", params, err)
	}

	fmt.Printf("  ============== Jobs [%d] =============\n", len(jobs.Jobs))
	fmt.Printf(" time %v user %+v\n", jobs.ServerTime, jobs.AuthUser)
	for _, job := range jobs.Jobs {
		fmt.Printf("~ %s - %s\n", job.Id, job.Title)
	}
}

// GetConfig for upwork client
func getConfig() api.ApiClient {
	cfg := api.ReadConfig(cfgFile)
	client := api.Setup(cfg)
	if !client.HasAccessToken() {
		getAccessToken()
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

func getHttpClient() {
	c := &http.Client{}
	config := api.ReadConfig(cfgFile)
	config.SetCustomHttpClient(c)
	client = api.Setup(cfg)
}

func getAccessToken() {
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
	err := saveConfig(cfgFile, cfg)
	if err != nil {
		log.Fatalf("failed to save config file %v", err)
	}
}
