package authentity

import (
	"time"

	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
)

type Resource struct {
	ResourceId  int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      commonentity.Status
}

type ResourceName string

const (
	ResourceAuth        ResourceName = "auth"
	ResourceUser        ResourceName = "user"
	ResourceUserProfile ResourceName = "user:profile"
)
