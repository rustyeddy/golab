package upw

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rustyeddy/golang-upwork/api/routers/auth"
	"github.com/rustyeddy/golang-upwork/api/routers/jobs/search"
)

// JobList is the result of the GetJobs request
type Jobs struct {
	ServerTime    float64 `json:"server_time"`
	ProfileAccess string  `json:"profile_access"`
	Jobs          []Job   `json:"jobs"`
	AuthUser      `json:"auth_user"`
}

// Job is an advertised paying gig
type Job struct {

	// Identities ...
	Id      string `json:"id"`
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	URL     string `json:"url"`

	Category2    string `json:"category2"`
	Subcategory2 string `json:"subcategory2"`

	Skills []string `json:"skills"`

	Budget      float64 `json:"budget"`
	Duration    string  `json:"duration"`
	Workload    string  `json:"workload"`
	JobStatus   string  `json:"job_status"`
	JobType     string  `json:"job_type"`
	DateCreated string  `json:"date_created"`

	Client `json:"client"`
}

// FetchJobs and Get Paid
func FetchJobs(params map[string]string) (*Jobs, error) {

	client := apiClient()

	/*
	 * Figure out who the user is and proceed with authorization
	 * http.Response and []byte will be return, you can use any
	 * TODO - need to respond properly to various results
	 */
	var resp *http.Response
	resp, _ = auth.New(client).GetUserInfo()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("FetchJobs: %d", resp.StatusCode)
	}

	var cont []byte
	upJobs := search.New(client)
	if resp, cont = upJobs.Find(params); resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to find a job %v", resp)
	}

	var jobs Jobs
	if err := json.Unmarshal(cont, &jobs); err != nil {
		return nil, fmt.Errorf("could not unmarshal %v", err)
	}
	return &jobs, nil
}
