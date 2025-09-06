package authutil

import "github.com/atareversei/quardian/services/api/internal/service/authservice"

var globalAuthService *authservice.Service

func Init(repo authservice.Repository) {
	svc := authservice.New(repo)
	globalAuthService = &svc
}

func AuthService() *authservice.Service {
	if globalAuthService == nil {
		panic("Auth service has not been initialized")
	}
	return globalAuthService
}
