package userdto

import (
	"time"

	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
	"github.com/atareversei/quardian/services/api/pkg/patch"
)

type ProfileRequest struct{}

type ProfileResponse struct {
	UserId     int                 `json:"user_id"`
	EmployeeId *string             `json:"employee_id"`
	FirstName  *string             `json:"first_name"`
	LastName   *string             `json:"last_name"`
	Username   string              `json:"username"`
	Email      *string             `json:"email"`
	Mobile     *string             `json:"mobile"`
	BirthDate  *time.Time          `json:"birth_date"`
	Status     commonentity.Status `json:"status"`
	Roles      []RawRole           `json:"roles"`
}

type RawRole struct {
	Id          int             `json:"id"`
	Permissions []RawPermission `json:"permissions"`
}

type RawPermission struct {
	Id      int   `json:"id"`
	Actions []int `json:"actions"`
}

type EditProfileRequest struct {
	FirstName patch.Null[string] `json:"first_name"`
	LastName  patch.Null[string] `json:"last_name"`
	Username  patch.Null[string] `json:"username"`
	Email     patch.Null[string] `json:"email"`
	Mobile    patch.Null[string] `json:"mobile"`
	BirthDate patch.Null[string] `json:"birth_date"`
}

type EditProfileResponse struct {
	ProfileResponse
}
