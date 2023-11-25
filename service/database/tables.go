package database

import "time"

type User struct {
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Biography string `json:"bio"`
	Token     string `json:"token"`
}

type Post struct {
	Pid       string    `json:"pid"`
	Uid       string    `json:"uid"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"createdAt"`
}

type Comment struct {
	Cid     string `json:"cid"`
	Pid     string `json:"pid"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
}
