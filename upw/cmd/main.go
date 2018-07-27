package main

import (
	"flag"
	"fmt"

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

	// Parse command line flags to know what we want to do
	flag.Parse()

	params := fromRequest(req)
	if jobs, err = upw.FetchJobs(params); err != nil {
		log.Fatal("FetchJobs params: ", params, err)
	}

	fmt.Printf("  ============== Jobs [%d] =============\n", len(jobs.Jobs))
	fmt.Printf(" time %v user %+v\n", jobs.ServerTime, jobs.AuthUser)
	for _, job := range jobs.Jobs {
		fmt.Printf("~ %s - %s\n", job.Id, job.Title)
	}
}

func fromRequest(req request) map[string]string {
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
	return params
}
