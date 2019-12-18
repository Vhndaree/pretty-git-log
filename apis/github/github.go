package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/vhndaree/pretty-git-log/interfaces"
	"github.com/vhndaree/pretty-git-log/util"
)

// Args - arguments to CallFunc function
type Args struct {
	TYPE string
	URL  string
}

type pullRequest = interfaces.PullRequest
type pullRequests = interfaces.PullRequests
type commits = interfaces.Commits
type user = interfaces.User

// CallFunc - calls function
func CallFunc(a Args) (res []byte) {
	switch a.TYPE {
	case "PULLREQUESTS":
		res = fetchPullRequests()
		break
	case "COMMITS_ON_PR":
		res = fetchCommitsOnPR(a.URL)
		break
	}

	return
}

func fetchPullRequests() []byte {
	n, d := GetUserInfo().UserName, util.GetTodaysDate()
	url := fmt.Sprintf("https://api.github.com/search/issues?q=author:%s+is:pr+created:2019-12-01..%s", n, d)

	return fetchFromGithub(url)
}

// fetchMyCommitsOnPR returns particular authors commit in PR
func fetchCommitsOnPR(url string) []byte {
	url = fmt.Sprintf("%s/commits", url)

	return fetchFromGithub(url)
}

// GetUserInfo - gives user belongs to token
func GetUserInfo() (u user) {
	url := fmt.Sprintf("https://api.github.com/user")
	resp := fetchFromGithub(url)
	err := json.Unmarshal(resp, &u)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func fetchFromGithub(url string) (body []byte) {
	myGithubAuthorization := "token " + os.Getenv("GITHUB_TOKEN")
	resp := util.GetWithAuthorization(url, myGithubAuthorization)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
