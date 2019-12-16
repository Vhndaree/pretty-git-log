package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Vhndaree/task-monitor/github/interfaces"
	"github.com/Vhndaree/task-monitor/util"
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
	myGithubUserName := os.Getenv("GITHUB_USERNAME")
	pullRequests = fetchPullRequests()

	for _, v := range pullRequests.PullRequests {
		path := v.PullRequest.URL
		commits = fetchMyCommitsOnPR(path, myGithubUserName)
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
	myGithubUserName := os.Getenv("GITHUB_USERNAME")
	url := fmt.Sprintf("https://api.github.com/search/issues?q=author:%s+is:pr+created:2019-11-29..%s", myGithubUserName, util.GetTodaysDate())
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
	url := fmt.Sprintf("%s/commits", path)
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
	resp := util.GetWithAuthorization(url, myGithubToken)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return body, nil
}
