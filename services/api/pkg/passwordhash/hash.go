package passwordhash

import (
	"github.com/atareversei/quardian/services/api/pkg/richerror"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	const op = "passwordhash.Hash"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", richerror.New(op).WithKind(richerror.KindUnexpected).WithError(err)
	}
	return string(hash), nil
}

func Compare(hashedPassword, plainPassword string) bool {
	const op = "passwordHash.Compare"
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
