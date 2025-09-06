package authdto

type LoginRequest struct {
	Username string `json:"username" example:"johndoe" minLength:"6" maxLength:"100"`
	Password string `json:"password" example:"xF54sal-M sa12" minLength:"8"`
}

type LoginResponse struct {
	Token string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
