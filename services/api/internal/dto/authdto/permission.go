package authdto

type IsUserIdPermittedOnResourceAndActionRequest struct {
	UserId     int
	ResourceId int
	ActionId   int
}

type IsUserIdPermittedOnResourceAndActionResponse struct {
	IsPermitted bool
}
