package events

type BusinessRegister struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"pass"`
	Role      string `json:"role"`
}
