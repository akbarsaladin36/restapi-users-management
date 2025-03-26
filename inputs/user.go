package inputs

type CreateUserInput struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"userRole"`
}

type UpdateUserInput struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"userRole"`
}
