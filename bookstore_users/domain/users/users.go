package users

type User struct {
	Id          int64  `json:"id"`
	FisrtName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}
