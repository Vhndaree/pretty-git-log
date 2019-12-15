package interfaces

// PullRequests - struct for PRs
type PullRequests struct {
	PullRequests []*PullRequest `json:"items"`
}

// PRs - asd sa
type PRs []PullRequest

// PullRequest - struct for PR
type PullRequest struct {
	PullNumber int    `json:"number"`
	Title      string `json:"title"`
	Link       string `json:"html_url"`
	URL        string `json:"url"`
	Body       string `json:"body"`
}
