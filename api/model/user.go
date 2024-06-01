package model

type User struct {
	Id         int    `json:"id" `
	User_Id    int    `json:"user_id" param:"user_id"`
	Faculty    string `json:"faculty"`
	Department string `json:"department"`
	Major      string `json:"major"`
	Course     string `json:"course"`
	Grade      int    `json:"grade"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
