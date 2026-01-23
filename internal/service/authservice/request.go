package authservice

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email         string
	Login         string
	Password      string
	CheckPassword string
}
