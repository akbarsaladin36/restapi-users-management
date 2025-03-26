package inputs

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
