package main

import (
	"bytes"
	"encoding/json"
	"log"

	"golang.org/x/text/search"
)

// JobList is the result of the GetJobs request
type JobList struct {
	ServerTime    int    `json:"server_time"`
	ProfileAccess string `json:"profile_access"`
	AuthUser      `json:"auth_user"`
	Jobs          []Job `json:"jobs"`
}

// Job is an advertised paying gig
type Job struct {
	Id      string
	Title   string
	Snippet string
	URL     string

	Category2    string
	Subcategory2 string

	// TODO turn skills into a type
	Skills []string

	JobType     string  `json:"job_type"`
	Budget      float64 `json:"budget"`
	Duration    string  `json:"duration"`
	Workload    string  `json:"workload"`
	JobStatus   string  `json:"job_status"`
	DateCreated string  `json:"date_Created"`

	Client
}

// AuthUser represents an Upwork account
type AuthUser struct {
	First          string `json:"first_name"`
	Last           string `json:"last_name"`
	Timezone       string `json:"timezone"`
	TimezoneOffset string `json:"timezone_offset"`
}

// Client is the person advertising the job looking to hire the client
type Client struct {
	Country      string
	Feedback     int
	ReviewsCount int `json:"reviews_count"`
	JobsPosted   int `json:"jobs_posted"`
	PastHires    int `json:"past_hires"`
}

// Get the Job List from Upwork based on search parameters
func getJobList() (*JobList, error) {

	params := make(map[string]string)
	if *title != "" {
		params["title"] = *title
	}
	if *query != "" {
		params["query"] = *query
	}
	if *skill != "" {
		params["skills"] = *skill
	}

	var jlist *JobList
	jlist, err = GetJobList(params)
	if err != nil {
		return nil, err
	}

	return jlist, nil
}

// GetJobList
func GetJobList(params map[string]string) (*JobList, error) {

	// Getting the jobs data
	// Get upwork job search client
	jobs := search.New(*UpClient)
	resp, data := jobs.Find(params)
	if resp.StatusCode != 200 {
		log.Fatalf("failed to query jobs - status %d - %b", resp.Status, data)
	}

	// validate the jobs data
	if !json.Valid(data) {
		log.Fatalf("json data appears to be invalid")
	}

	// Indent the json and write it to STDOUT or a file if requested
	var jindented bytes.Buffer
	json.Indent(&jindented, data, "", "\t")

	// jindented.WriteTo(os.Stdout)

	var jlist JobList
	err := json.Unmarshal(data, &jlist)
	if err != nil {
		log.Fatalf("failed to marshal json %v", err)
	}
	return &jlist, nil
}
