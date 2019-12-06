package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchRequest struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreateAt time.Time `json:"Created_at"`
	Body string // Markdown格式
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}
