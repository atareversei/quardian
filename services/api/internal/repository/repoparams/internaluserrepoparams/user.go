package internaluserrepoparams

import (
	"time"

	"github.com/atareversei/quardian/services/api/internal/dto"
	"github.com/atareversei/quardian/services/api/internal/dto/internaluserdto"
)

type ListUsersInput struct {
	Filters internaluserdto.UserFilters
}

type ListUsersOutputItem struct {
	internaluserdto.ListUsersResponseItem
	CreatedAtRepo time.Time
}

type ListUsersOutput struct {
	List []ListUsersOutputItem
	Meta dto.ListMeta
}

func FromListUsersRepoParamToDTO(list []ListUsersOutputItem) []internaluserdto.ListUsersResponseItem {
	dtos := make([]internaluserdto.ListUsersResponseItem, len(list))
	for i, p := range list {
		dtos[i] = p.ListUsersResponseItem
	}
	return dtos
}
