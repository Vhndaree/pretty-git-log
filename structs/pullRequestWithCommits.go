package structs

// PullRequestsWithCommits - structure for PRs with its commits
type PullRequestsWithCommits []PullRequestWithCommits

// PullRequestWithCommits - structure for PR with its commits
type PullRequestWithCommits struct {
	PullRequest PullRequest
	Commits     Commits
}
