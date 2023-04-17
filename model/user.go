package model

type User struct {
	ID        int
	Login     string `json:"login" form:"login"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Password  string `json:"password" form:"password"`
	Status	string	`json:"status" form:"status"`
}

type SignInData struct {
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
}
