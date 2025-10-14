package configdto

import (
	"github.com/atareversei/quardian/services/api/internal/dto"
	"github.com/atareversei/quardian/services/api/internal/entity/commonentity"
	"net"
)

type DeviceFilters struct {
	Page    int
	PerPage int
	Status  commonentity.Status
}

type ListDevicesRequest struct {
	DeviceFilters
}

type ListDevicesResponseItem struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	IP          net.IP              `json:"ip"`
	Status      commonentity.Status `json:"status"`
}

type ListDevicesResponse struct {
	Meta dto.ListMeta              `json:"meta"`
	List []ListDevicesResponseItem `json:"devices"`
}
