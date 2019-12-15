package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"task-monitor/github/interfaces"
	"task-monitor/service"
)

type pullRequestsWithCommits = interfaces.PullRequestsWithCommits
type pullRequestWithCommits = interfaces.PullRequestWithCommits
type pullRequest = interfaces.PullRequest
type pullRequests = interfaces.PullRequests
type commits = interfaces.Commits

// FetchPullRequestsWithCommits - gives PR with its commits
func FetchPullRequestsWithCommits() pullRequestsWithCommits {
	var commits commits
	var pullRequests pullRequests
	var pullRequestWithCommits pullRequestWithCommits
	var pullRequestsWithCommits pullRequestsWithCommits
	myUseName := "Vhndaree"
	pullRequests = fetchPullRequests()

	for _, v := range pullRequests.PullRequests {
		path := v.PullRequest.URL
		commits = fetchMyCommitsOnPR(path, myUseName)
		pullRequestWithCommits.PullRequest = v
		pullRequestWithCommits.Commits = commits

		pullRequestsWithCommits = append(pullRequestsWithCommits, pullRequestWithCommits)
	}

	return pullRequestsWithCommits
}

// fetchPullRequests - gives PR info
func fetchPullRequests() pullRequests {
	var PRResponse pullRequests
	var emptyResponse pullRequests
	url := "https://api.github.com/search/issues?q=author:Vhndaree+is:pr+created:2019-08-01.." + service.GetTodaysDate()
	resp, err := fetchFromGithub(url)

	if err != nil {
		log.Fatal(err)
		return emptyResponse
	}

	err = json.Unmarshal(resp, &PRResponse)
	if err != nil {
		log.Fatal(err)
		return emptyResponse
	}

	return PRResponse
}

// fetchMyCommitsOnPR returns particular authors commit in PR
func fetchMyCommitsOnPR(path, myUserName string) commits {
	var commitResponse commits
	var emptyResponse commits
	url := path + "/commits"
	resp, err := fetchFromGithub(url)
	if err != nil {
		log.Fatal(err)
		return emptyResponse
	}

	err = json.Unmarshal(resp, &commitResponse)
	if err != nil {
		log.Fatal(err)
		return emptyResponse
	}

	return filterCommitsByUser(commitResponse, myUserName)
}

// filterCommitsByUser - filters commits by username
func filterCommitsByUser(c commits, userName string) commits {
	var temp commits

	for _, v := range c {
		if strings.ToLower(v.UserName.Login) == strings.ToLower(userName) {
			temp = append(temp, v)
		}
	}

	return temp
}

func fetchFromGithub(url string) ([]byte, error) {
	myGithubToken := "token " + os.Getenv("GITHUB_TOKEN")
	resp := service.GetWithAuthorization(url, myGithubToken)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return body, nil
}
