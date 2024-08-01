package schema

type LoginRequestBody struct {
	Username string `json:"username"`
}

type UserLogin struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
