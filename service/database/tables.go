package database

import "time"

type User struct {
	Uid       string
	Username  string
	Biography string
	Token     string
}

type Post struct {
	Pid       string
	Uid       string
	Caption   string
	CreatedAt time.Time
}

type Comment struct {
	Cid     string
	Pid     string
	Uid     string
	content string
}
