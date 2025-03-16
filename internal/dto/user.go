package dto

type UserDataResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Created  string `json:"created"`
}
