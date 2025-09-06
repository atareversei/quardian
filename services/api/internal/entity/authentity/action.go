package authentity

import (
	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
	"time"
)

type Action struct {
	ActionId    int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      commonentity.Status
}

type ActionName string

const (
	ActionRead ActionName = "read"
)
