package model

type Login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func NewLogin() *Login {
	return &Login{}
}
