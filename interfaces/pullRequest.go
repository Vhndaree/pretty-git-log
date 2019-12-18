package interfaces

// PullRequests - struct for PRs
type PullRequests []PullRequest

// PullRequest - struct for PR
type PullRequest struct {
	PullNumber float64
	Title      string
	Link       string
	Body       string
	URL        string
}

type prInfo struct {
	URL string
}
