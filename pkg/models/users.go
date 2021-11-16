package models

type User struct {
	ID            int      `json:"-"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	StudentNumber string      `json:"student_number"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterUser struct{
	Username string `json:"username"`

}
