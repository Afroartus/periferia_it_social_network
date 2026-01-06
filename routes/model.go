package routes

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreatePostRequest struct {
	Message string `json:"message"`
}

type UserRequest struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}
type ProfileResponse struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
