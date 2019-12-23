package github

import (
	"encoding/json"
	"log"

	"github.com/Vhndaree/pretty-git-log/apis/github"
	"github.com/Vhndaree/pretty-git-log/structs"
)

type pullRequestsWithCommits = structs.PullRequestsWithCommits
type pullRequestWithCommits = structs.PullRequestWithCommits
type pullRequest = structs.PullRequest
type pullRequests = structs.PullRequests
type commits = structs.Commits
type args = github.Args

type githubPullRequests pullRequests

func (gh githubPullRequests) getPullRequests() (prs pullRequests) {
	arg := args{TYPE: "PULLREQUESTS", URL: ""} //we should add another flag vc for recognize github or gitlab .... on future version
	resp := github.CallFunc(arg)

	var r map[string]interface{}
	err := json.Unmarshal(resp, &r)
	if err != nil {
		log.Fatal(err)
		return
	}

	items := r["items"].([]interface{})
	for _, item := range items {
		i := item.(map[string]interface{})
		info := i["pull_request"].(map[string]interface{})
		pr := pullRequest{
			PullNumber: i["number"].(float64),
			Title:      i["title"].(string),
			Link:       i["html_url"].(string),
			Body:       i["body"].(string),
			URL:        info["url"].(string),
		}
		prs = append(prs, pr)
	}

	return
}
