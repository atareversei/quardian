package authdto

type RegisterRequest struct {
	Username string `json:"username" example:"johndoe" minLength:"6" maxLength:"100"`
	Password string `json:"password" example:"xF54sal-M sa12" minLength:"8"`
}

type RegisterResponse struct {
}
