package github

import (
	"encoding/json"
	"log"

	"github.com/Vhndaree/pretty-git-log/apis/github"
	"github.com/Vhndaree/pretty-git-log/interfaces"
	"github.com/Vhndaree/pretty-git-log/util"
)

type pullRequestsWithCommits = interfaces.PullRequestsWithCommits
type pullRequestWithCommits = interfaces.PullRequestWithCommits
type pullRequest = interfaces.PullRequest
type pullRequests = interfaces.PullRequests
type commits = interfaces.Commits
type args = github.Args

// FetchMyPullRequestsWithMyCommits - gives PR with its commits
func FetchMyPullRequestsWithMyCommits() (r pullRequestsWithCommits) {
	var pullRequests pullRequests
	var pullRequestWithCommits pullRequestWithCommits
	pullRequests = getMyPullRequests()
	u := github.GetUserInfo()

	for _, v := range pullRequests {
		url := v.URL
		c := util.FilterCommitsByUser(getCommitsOnPR(url), u.UserName)
		pullRequestWithCommits.PullRequest = v
		pullRequestWithCommits.Commits = c

		r = append(r, pullRequestWithCommits)
	}

	return
}

// getMyPullRequests - gives pull requests belongs to GITHUB_TOKEN
func getMyPullRequests() (prs pullRequests) {
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

func getCommitsOnPR(url string) (commits commits) {
	arg := args{TYPE: "COMMITS_ON_PR", URL: url}
	resp := github.CallFunc(arg)
	err := json.Unmarshal(resp, &commits)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
