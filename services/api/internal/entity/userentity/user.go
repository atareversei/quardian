package userentity

import (
	"time"

	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

// ? Split into two entities - one for the base user and one for profile
type User struct {
	UserId     int
	UserName   string
	EmployeeId string
	FirstName  string
	LastName   string
	Email      string
	Mobile     string
	Address    string
	BirthDate  time.Time
	CreatedAt  time.Time
	Status     commonentity.Status
}

type UserWithPasswordHash struct {
	User
	PasswordHash string
}
