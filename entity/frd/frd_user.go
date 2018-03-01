package frd

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type UserProfile struct {
	UserProfile User `json:"profile"`
}
