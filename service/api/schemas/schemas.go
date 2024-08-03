package schema

import "errors"

var ErrExistingUsername = errors.New("usernme already exists")

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
