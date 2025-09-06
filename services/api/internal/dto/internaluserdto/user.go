package internaluserdto

import (
	"github.com/atareversei/quardian/services/api/internal/dto"
	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

type UserFilters struct {
	Page    int
	PerPage int
	Status  commonentity.Status
}

type ListUsersRequest struct {
	UserFilters
}

type ListUsersResponseItem struct {
	UserId     int                   `json:"user_id"`
	UserName   string                `json:"username"`
	EmployeeId *string               `json:"employee_id"`
	FirstName  *string               `json:"first_name"`
	LastName   *string               `json:"last_name"`
	CreatedAt  commonentity.DateTime `json:"created_at"`
	Status     commonentity.Status   `json:"status"`
}

type ListUsersResponse struct {
	Meta dto.ListMeta            `json:"meta"`
	List []ListUsersResponseItem `json:"users"`
}
