package interfaces

// PullRequests - struct for PRs
type PullRequests struct {
	PullRequests []*PullRequest `json:"items"`
}

// PullRequest - struct for PR
type PullRequest struct {
	PullNumber  int    `json:"number"`
	Title       string `json:"title"`
	Link        string `json:"html_url"`
	Body        string `json:"body"`
	PullRequest prInfo `json:"pull_request"`
}

type prInfo struct {
	URL string `json:"url"`
}
