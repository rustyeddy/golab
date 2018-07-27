package upw

// JobList is the result of the GetJobs request
type Jobs struct {
	ServerTime    int    `json:"server_time"`
	ProfileAccess string `json:"profile_access"`
	AuthUser      `json:"auth_user"`
	Jobs          []Job `json:"jobs"`
}

// Job is an advertised paying gig
type Job struct {
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
