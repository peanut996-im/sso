package handler

type login struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func newLogin() *login {
	return &login{}
}
