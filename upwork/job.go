package main

type Job struct {
}

func NewJob(info map[string]interface{}) *Job {
	return &Job{info}
}

func (j *Job) Budget() int {
	return j.info["budget"]
}

// JobList is the result of the GetJobs request
type JobList struct {
	ServerTime    int    `json:"server_time"`
	ProfileAccess string `json:"profile_access"`
	AuthUser      `json:"auth_user"`
	Jobs          []Job `json:"jobs"`
}

// Job is an advertised paying gig
type Job struct {
	response map[string]interface{}
	Id       string
	Title    string
	Snippet  string
	URL      string

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
