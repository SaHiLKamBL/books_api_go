package types

type UserRes struct{
	Name string `json:"name"`
	Email  string `json:"email"`
}

type LoginUserReq struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	AccessToken string `json:"access_token"`
	User  UserRes `json:"user"`
}