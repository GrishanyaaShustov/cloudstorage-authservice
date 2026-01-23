package authservice

type LoginResponse struct {
	AccessToken        string
	RefreshTokenID     string
	RefreshTokenSecret string
}

type RegisterResponse struct {
	AccessToken        string
	RefreshTokenID     string
	RefreshTokenSecret string
}
