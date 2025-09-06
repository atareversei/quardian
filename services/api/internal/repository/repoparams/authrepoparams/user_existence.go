package authrepoparams

type DoesUserNameWithPasswordExistInput struct {
	UserName     string
	PasswordHash string
}

type DoesUserNmeExistInput struct {
	UserName string
}
