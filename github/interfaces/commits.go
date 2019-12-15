package interfaces

import "time"

// Commits holds commit log
type Commits []*commit

type commit struct {
	Hash          string        `json:"sha"`
	URL           string        `json:"html_url"`
	CommitDetails commitDetails `json:"commit"`
	UserName      author        `json:"author"`
}

type commitDetails struct {
	Committer committer `json:"author"`
	Message   string    `json:"message"`
}

type author struct {
	Login string `json:"login"`
}

type committer struct {
	Date *time.Time `json:"date"`
}
