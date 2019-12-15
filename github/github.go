package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Vhndaree/task-monitor/github/interfaces"
	"github.com/Vhndaree/task-monitor/service"
)

type pullRequestsWithCommits = interfaces.PullRequestsWithCommits
type pullRequestWithCommits = interfaces.PullRequestWithCommits
type pullRequest = interfaces.PullRequest
type prs = interfaces.PRs
type pullRequests = interfaces.PullRequests
type commits = interfaces.Commits

type repositories []repository

type repository struct {
	Owner        string
	Repositories []string
}

var myUserName = "Vhndaree"
var sinceDate = "2019-08-01"

// CombineAllPRs - returns list of own and others PR with own commits
func CombineAllPRs() pullRequestsWithCommits {
	myPRs := fetchMyPullRequestsWithCommits()
	// myCommitOnOthersPRs := fetchMyCommitsAndItsPR()
	combinedPRs := myPRs

	// for _, v := range myCommitOnOthersPRs {
	// 	combinedPRs = append(combinedPRs, v)
	// }

	return combinedPRs
}

// fetchMyPullRequestsWithCommits - gives PR with its commits
func fetchMyPullRequestsWithCommits() pullRequestsWithCommits {
	var commits commits
	var pullRequests pullRequests
	var pullRequestWithCommits pullRequestWithCommits
	var pullRequestsWithCommits pullRequestsWithCommits
	pullRequests = fetchPullRequests()

	for _, v := range pullRequests.PullRequests {
		path := service.Before(v.URL, "/issues") + "/pulls/" + strconv.Itoa(v.PullNumber)
		commits = fetchMyCommitsOnMyPR(path, myUserName)
		pullRequestWithCommits.PullRequest = v
		pullRequestWithCommits.Commits = commits

		pullRequestsWithCommits = append(pullRequestsWithCommits, pullRequestWithCommits)
	}
	return pullRequestsWithCommits
}

// fetchMyCommitsAndItsPR - fetch commits and map with its respective PR
func fetchMyCommitsAndItsPR() pullRequestsWithCommits {
	var allCommits commits
	var prResponse prs
	var prWithCommit pullRequestWithCommits
	var prsWithCommit pullRequestsWithCommits
	var emptyResponse pullRequestsWithCommits
	allCommits = fetchMyCommitsOnMaster()

	for _, v := range allCommits {
		url := v.APIURL + "/pulls"
		resp, err := fetchFromGithub(url)
		if err != nil {
			log.Fatal(err)
			return emptyResponse
		}

		err = json.Unmarshal(resp, &prResponse)
		if err != nil {
			log.Fatal(err)
			return emptyResponse
		}
		for _, pr := range prResponse {
			prWithCommit.PullRequest = &pr
		}
		prWithCommit.Commits = append(prWithCommit.Commits, v)
		prsWithCommit = append(prsWithCommit, prWithCommit)
	}

	if len(prsWithCommit) < 1 {
		return emptyResponse
	}

	return combineCommitsOfSamePR(prsWithCommit)
}

// fetchPullRequests - gives PR info
func fetchPullRequests() pullRequests {
	var PRResponse pullRequests
	var emptyResponse pullRequests
	url := "https://api.github.com/search/issues?q=author:Vhndaree+is:pr+created:" + sinceDate + ".." + service.GetTodaysDate()
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

// fetchMyCommitsOnMyPR returns particular authors commit in PR
func fetchMyCommitsOnMyPR(path, myUserName string) commits {
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

// fetchMyCommitsOnMaster - returns Commits
func fetchMyCommitsOnMaster() commits {
	var repositories repositories
	var repository repository
	var allCommits commits
	repository.Owner = "phil-inc"
	repository.Repositories = []string{"capi", "plib", "phil-me", "aapi", "phil-md-dashboard", "pp-dash", "phil-web"}
	repositories = append(repositories, repository)

	for _, v := range repositories {
		for _, r := range v.Repositories {
			var commitsOnRepo commits
			var emptyResponse commits
			url := "https://api.github.com/repos/" + v.Owner + "/" + r + "/commits?author=" + myUserName + "&since=" + sinceDate
			resp, err := fetchFromGithub(url)
			if err != nil {
				log.Fatal(err)
				return emptyResponse
			}

			err = json.Unmarshal(resp, &commitsOnRepo)
			if err != nil {
				log.Fatal(err)
				return emptyResponse
			}

			for _, i := range commitsOnRepo {
				allCommits = append(allCommits, i)
			}
		}
	}

	return allCommits
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

func combineCommitsOfSamePR(prsWithCommit pullRequestsWithCommits) pullRequestsWithCommits {
	var prLinks []string
	var combinedPRsWithCommits pullRequestsWithCommits

	for i := 0; i < len(prsWithCommit); i++ {
		if service.Contains(prLinks, prsWithCommit[i].PullRequest.Link) {
			continue
		}

		var prWithCommits pullRequestWithCommits
		prWithCommits.PullRequest = prsWithCommit[i].PullRequest
		for j := i + 1; j < len(prsWithCommit); j++ {

			if prsWithCommit[i].PullRequest.Link == prsWithCommit[j].PullRequest.Link {
				prsWithCommit[i].Commits = append(prsWithCommit[i].Commits, prsWithCommit[j].Commits[0])
			}
		}
		prWithCommits.Commits = prsWithCommit[i].Commits
		combinedPRsWithCommits = append(combinedPRsWithCommits, prWithCommits)
		prLinks = append(prLinks, prsWithCommit[i].PullRequest.Link)
	}

	return combinedPRsWithCommits
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
