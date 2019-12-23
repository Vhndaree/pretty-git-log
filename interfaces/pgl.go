package interfaces

import "github.com/Vhndaree/pretty-git-log/structs"

type pullRequestsWithCommits = structs.PullRequestsWithCommits
type pullRequests = structs.PullRequests
type commits = structs.Commits

type prettyGitLog interface {
	getPullRequestsWithCommits() pullRequestsWithCommits
	getPullRequests() pullRequests
	getCommitsOnPR(url string) commits
}
