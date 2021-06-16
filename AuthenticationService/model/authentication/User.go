package authentication

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Role string `json:"role"`
}