package database

import "time"

type User struct {
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Biography string `json:"bio"`
	Token     string `json:"token"`
}

type UserProfile struct {
	User         User   `json:"user"`
	Posts        []Post `json:"posts"`
	ProfileImage string `json:"profileImage"`
	Follower     int    `json:"follower"`
	Follows      int    `json:"follows"`
}

type Post struct {
	Pid       string    `json:"pid"`
	Uid       string    `json:"uid"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"createdAt"`
	Image     string    `json:"image"`
}

type Comment struct {
	Cid     string `json:"cid"`
	Pid     string `json:"pid"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
}
