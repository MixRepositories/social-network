package structs

type User struct {
	Id        uint16 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Password  string `json:"password"`
	Birthday  string `json:"birthday"`
}

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
