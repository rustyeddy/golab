package upw

// AuthUser represents an Upwork account
type AuthUser struct {
	First          string `json:"first_name"`
	Last           string `json:"last_name"`
	Timezone       string `json:"timezone"`
	TimezoneOffset string `json:"timezone_offset"`
}

// Client is the person advertising the job looking to hire the client
type Client struct {
	Country      string  `json:"country"`
	Feedback     float64 `json:"feedback"`
	ReviewsCount int     `json:"reviews_count"`
	JobsPosted   int     `json:"jobs_posted"`
	PastHires    int     `json:"past_hires"`
}
