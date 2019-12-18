package file

import (
	"fmt"
	"log"
	"os"

	"github.com/vhndaree/pretty-git-log/interfaces"
	"github.com/vhndaree/pretty-git-log/service/github"
	"github.com/vhndaree/pretty-git-log/util"
)

var getSpaces = util.GetSpacesOfLength

// Write - create and write into file
func Write() {
	// create file for writting
	file, err := os.Create("/home/lf/Downloads/for_timesheet/Vhndaree_" + util.GetTodaysDate() + ".text")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	content := getContent()
	fmt.Fprint(file, content)

	file.Close()
	fmt.Print("File created successfully.....")
}

func getContent() string {
	content := ""
	var ghr interfaces.PullRequestsWithCommits
	ghr = github.FetchMyPullRequestsWithMyCommits()
	for _, v := range ghr {
		pn := fmt.Sprintf("%d", int(v.PullRequest.PullNumber))
		spaces := getSpaces(len(pn))
		content += "#" + pn + "\t" + "PR: " + v.PullRequest.Link + "\n"
		content += spaces + v.PullRequest.Title + "\n"
		content += spaces + "Description: " + v.PullRequest.Body + "\n\n"
		content += spaces + "----------Commits----------"

		count := 1
		LastCommitDate := "2000-01-01"
		for _, commit := range v.Commits {

			if LastCommitDate != commit.CommitDetails.Committer.Date.Format("2006-01-02") {
				content += "\n\n" + spaces + "Date: " + commit.CommitDetails.Committer.Date.Format("2006-01-02") + ",\n"
			}

			LastCommitDate = commit.CommitDetails.Committer.Date.Format("2006-01-02")
			content += spaces + spaces + "SHA: " + commit.Hash + "\n"
			content += spaces + spaces + "URL: " + commit.URL + "\n"
			content += spaces + spaces + "Message: " + commit.CommitDetails.Message + "\n"
			if count < len(v.Commits) {
				content += "\n\n"
			} else {
				content += "\n"
			}

			count++
		}
		content += spaces + "--------------------------------------------------------------------------------\n\n\n"
	}

	return content
}
