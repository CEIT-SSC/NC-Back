package models

type User struct {
	ID       int `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	StudentNumber int `json:"student_number"`
	Tokens []string `json:"-"`
}