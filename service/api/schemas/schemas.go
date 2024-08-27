package schema

import "errors"

var ErrExistingUsername = errors.New("usernme already exists")
var ErrNoAuthentication = errors.New("not Authenticated")
var ErrNotAuthorized = errors.New("not Authorized")

type LoginRequestBody struct {
	Username string `json:"username"`
}

type UserLogin struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SimpleUserData struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

type FollowData struct {
	FollowingId string `json:"followingId"`
	FollowedId  string `json:"followedId"`
}

type BanData struct {
	BanningId string `json:"banningId"`
	BannedId  string `json:"bannedId"`
}

type SimplePostData struct {
	Pid       string         `json:"pid"`
	ImageUrl  string         `json:"imageUrl"`
	CreatedAt string         `json:"uploadTime"`
	Author    SimpleUserData `json:"author"`
}

type UserProfileData struct {
	User  SimpleUserData   `json:"user"`
	Posts []SimplePostData `json:"posts"`
}

type LikeData struct {
	Uid string `json:"uid"`
	Pid string `json:"pid"`
}

type PostComment struct {
	Cid    string         `json:"cid"`
	Author SimpleUserData `json:"author"`
	Pid    string         `json:"pid"`
	Text   string         `json:"text"`
}

type PostData struct {
	Pid           string         `json:"pid"`
	ImageUrl      string         `json:"imageUrl"`
	Author        SimpleUserData `json:"author"`
	UploadTime    string         `json:"uploadTime"`
	LikesCount    int            `json:"likesCount"`
	CommentsCount int            `json:"commentsCOunt"`
	Comments      []PostComment  `json:"comments"`
}
