package authrepoparams

type CreateUserInput struct {
	UserName     string
	PasswordHash string
}

type CreateUserOutput struct {
	UserId int
}
